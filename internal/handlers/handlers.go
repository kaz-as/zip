package handlers

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"os"
	"path/filepath"
	"time"
	"unsafe"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kaz-as/zip/domain"
	"github.com/kaz-as/zip/models"
	"github.com/kaz-as/zip/pkg/archive"
	"github.com/kaz-as/zip/restapi/operations"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (s *HandlerSet) CheckChunksHandler(params operations.CheckChunksParams) middleware.Responder {
	chunks, err := s.ArchiveUseCaseMake(nil).GetUncompletedChunks(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return operations.NewCheckChunksNotFound()
		}
		return operations.NewCheckChunksDefault(0)
	}

	var payload []models.ChunkNumber

	for _, chunk := range chunks {
		payload = append(payload, models.ChunkNumber(chunk.Number))
	}

	return operations.NewCheckChunksOK().WithPayload(payload)
}

func (s *HandlerSet) CreateArchiveHandler(params operations.CreateArchiveParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	var filesForArchive []*archive.File

	retError := func() middleware.Responder {
		for _, fp := range filesForArchive {
			if fp == nil {
				continue
			}
			err := fp.In.Close()
			if err != nil {
				s.Log.Warn("file close: %s", err)
			}
		}
		return operations.NewCreateArchiveDefault(0)
	}

	uc := s.ArchiveUseCaseMake(nil)

	for _, a := range params.Files {
		inputFiles := a.Files
		archiveID := a.ID

		currArchive, err := uc.GetByID(ctx, int64(archiveID))
		if err != nil {
			s.Log.Info("s.ArchiveUseCaseMake.GetByID: %s", err)
			return retError()
		}

		for _, file := range inputFiles {
			path := []string{s.FolderForFiles, currArchive.Original}
			path = append(path, file.Path...)
			path = append(path, string(file.Name))

			finalFilePath := filepath.Join(path...)
			f, err := os.Open(finalFilePath)
			if err != nil {
				s.Log.Info("os.Open(\"%s\"): %s", finalFilePath, err)
				return retError()
			}

			filesForArchive = append(filesForArchive, &archive.File{
				In:          f,
				OutLocation: file.NewPath,
				OutName:     string(file.NewName),
			})
		}
	}

	zipOut, err := s.ArchiveService.Zip(ctx, filesForArchive...)
	if err != nil {
		s.Log.Info("zip: %s", err)
		return retError()
	}

	return operations.NewCreateArchiveOK().WithPayload(zipOut)
}

func (s *HandlerSet) GetFilesHandler(params operations.GetFilesParams) middleware.Responder {
	// todo: implement me
	return middleware.NotImplemented("operation GetFiles has not yet been implemented")
}

func (s *HandlerSet) InitUploadArchiveHandler(params operations.InitUploadArchiveParams) middleware.Responder {
	size := int64(*params.Archive.Size)
	name := string(*params.Archive.Name)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano() + size))
	rndName := randName(rnd)

	ctx := params.HTTPRequest.Context()

	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		s.Log.Error("begin tx: %s", err)
		return operations.NewInitUploadArchiveDefault(0)
	}
	defer s.rollback(tx)

	uc := s.ArchiveUseCaseMake(tx)

	original := rndName + name

	arch := domain.Archive{
		Original: original,
		Size:     size,
		Name:     name,
		InitTime: time.Now(),
	}

	chunks, err := uc.Store(params.HTTPRequest.Context(), &arch)
	if err != nil {
		s.Log.Error("store archive: %s", err)
		return operations.NewInitUploadArchiveDefault(0)
	}

	a, err := os.Create(s.FolderForArchives + original)
	if err != nil {
		s.Log.Error("create archive file: %s", err)
		return operations.NewInitUploadArchiveDefault(0)
	}
	defer func() {
		err := a.Close()
		if err != nil {
			s.Log.Error("cannot close file: %s", err)
		}
	}()

	err = a.Truncate(size)
	if err != nil {
		s.Log.Error("truncate archive file to size %v: %s", size, err)
		return operations.NewInitUploadArchiveDefault(0)
	}

	var eachSize int64
	if len(chunks) > 0 {
		eachSize = chunks[0].NextChunkByte - chunks[0].StartByte
	}

	ret := operations.NewInitUploadArchiveOK().WithPayload(&models.InitUploadSuccess{
		Chunks: models.ChunkNumber(len(chunks)),
		Each:   models.Size(eachSize),
		ID:     models.ArchiveID(arch.ID),
	})

	err = tx.Commit()
	if err != nil {
		s.Log.Error("commit failed: %s", err)
		return operations.NewInitUploadArchiveDefault(0)
	}

	return ret
}

func (s *HandlerSet) UploadChunkHandler(params operations.UploadChunkParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		s.Log.Error("begin tx: %s", err)
		return operations.NewInitUploadArchiveDefault(0)
	}
	defer s.rollback(tx)

	uc := s.ArchiveUseCaseMake(tx)

	// Upload chunk

	chunk, err := uc.GetChunkForUpdate(ctx, params.ID, params.QueryChunk)
	// Idempotency
	if chunk.Uploaded {
		err = tx.Commit()
		if err != nil {
			s.Log.Error("commit failed: %s", err)
			return operations.NewUploadChunkDefault(0)
		}
		return operations.NewUploadChunkOK()
	}

	err = s.ChunkWriterService.WriteChunk(
		params.Chunk,
		filepath.Join(s.FolderForArchives, chunk.Archive.Original),
		chunk.StartByte,
		chunk.NextChunkByte-chunk.StartByte,
	)
	if err != nil {
		s.Log.Error("write chunk to file failed: %s", err)
		return operations.NewUploadChunkDefault(0)
	}

	chunk.Uploaded = true
	chunk.UploadedTime = time.Now()
	err = uc.UpdateChunk(ctx, &chunk)
	if err != nil {
		s.Log.Error("update chunk failed: %s", err)
		return operations.NewUploadChunkDefault(0)
	}

	ret := operations.NewUploadChunkOK()
	err = tx.Commit()
	if err != nil {
		s.Log.Error("commit failed: %s", err)
		return operations.NewUploadChunkDefault(0)
	}

	// Unzip if all chunks are completed
	lastChunks, err := uc.GetUncompletedChunks(ctx, params.ID)
	if err == nil && len(lastChunks) == 0 {
		s.startUnzipping(chunk.Archive.ID)
	}

	return ret
}

func (s *HandlerSet) IsCompletedHandler(params operations.IsCompletedParams) middleware.Responder {
	uc := s.ArchiveUseCaseMake(nil)
	chunks, err := uc.GetUncompletedChunks(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return operations.NewIsCompletedNotFound()
		}
		return operations.NewIsCompletedDefault(0)
	}

	if len(chunks) > 0 {
		return operations.NewIsCompletedOK().WithPayload(false)
	}

	completed, err := uc.CheckCompleted(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return operations.NewIsCompletedNotFound()
		}
		return operations.NewIsCompletedDefault(0)
	}

	if !completed {
		s.startUnzipping(params.ID)
	}

	return operations.NewIsCompletedOK().WithPayload(completed)
}

func (s *HandlerSet) startUnzipping(id int64) {
	s.mx.Lock()
	defer s.mx.Unlock()

	arch, err := s.ArchiveUseCaseMake(nil).GetByID(context.Background(), id)
	if err != nil {
		s.Log.Error("get archive id=%v failed: %s", id, err)
		return
	}

	if arch.Unarchived {
		return
	}

	if _, ok := s.isUnzipping[arch.ID]; ok {
		return
	}
	s.isUnzipping[arch.ID] = struct{}{}

	go func(arch *domain.Archive) {
		defer func() {
			s.mx.Lock()
			defer s.mx.Unlock()

			delete(s.isUnzipping, arch.ID)
		}()

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
		defer cancel()

		name := s.FolderForArchives + arch.Original
		a, err := os.Open(name)
		if err != nil {
			s.Log.Error("cannot open archive id=%v: %s", arch.ID, err)
			return
		}
		defer func() {
			err := a.Close()
			if err != nil {
				s.Log.Error("file close failed: %s", err)
			}
		}()

		err = s.ArchiveService.Unzip(ctx, a, filepath.Join(s.FolderForFiles, arch.Original))
		if err != nil {
			s.Log.Error("cannot unzip archive id=%v: %s", arch.ID, err)
			return
		}

		s.Log.Info("archive id=%v is completed", arch.ID)

		err = s.ArchiveUseCaseMake(nil).SetCompleted(ctx, id, true)
		if err != nil {
			s.Log.Error("set completed failed archive id=%v: %s", id, err)
			return
		}
	}(&arch)
}

func (s *HandlerSet) rollback(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil && !errors.Is(err, sql.ErrTxDone) {
		s.Log.Error("tx rollback: %s", err)
	}
}

func randName(rnd *rand.Rand) string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = letters[rnd.Intn(len(letters))]
	}
	return *(*string)(unsafe.Pointer(&b))
}
