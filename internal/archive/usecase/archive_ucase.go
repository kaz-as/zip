package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/kaz-as/zip/domain"
)

const (
	defaultChunkSize = 1024 * 1024 * 10
	maxArchiveSize   = defaultChunkSize*(int64(math.MaxInt32)+1) - 1
)

var (
	ErrArchiveTooLarge = errors.New("the archive is too large")
	ErrNilTx           = errors.New("cannot proceed complicated updates outside of a transaction")
)

type archiveUseCase struct {
	archiveRepo domain.ArchiveRepository
	chunkRepo   domain.ChunkRepository
	ctxTimeout  time.Duration
	tx          *sql.Tx
}

func NewMaker(archiveRepo domain.ArchiveRepository, chunkRepo domain.ChunkRepository, timeout time.Duration) domain.MakeByTx {
	return func(tx *sql.Tx) domain.ArchiveUseCase {
		return &archiveUseCase{
			archiveRepo: archiveRepo,
			chunkRepo:   chunkRepo,
			ctxTimeout:  timeout,
		}
	}
}

func (a *archiveUseCase) GetByID(ctx context.Context, id int64) (domain.Archive, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	return a.archiveRepo.GetByID(ctxTimeout, a.tx, id)
}

func (a *archiveUseCase) GetChunkForUpdate(ctx context.Context, archiveID int64, chunkNumber int32) (domain.Chunk, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	chunk, err := a.chunkRepo.GetByNumber(ctxTimeout, a.tx, archiveID, chunkNumber)
	if err != nil {
		return chunk, err
	}

	err = a.attachArchiveToChunk(ctxTimeout, &chunk, archiveID)
	return chunk, err
}

func (a *archiveUseCase) UpdateChunk(ctx context.Context, chunk *domain.Chunk) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	return a.chunkRepo.Update(ctxTimeout, a.tx, chunk)
}

func (a *archiveUseCase) GetUncompletedChunks(ctx context.Context, archiveID int64) ([]domain.Chunk, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	chunks, err := a.chunkRepo.GetUncompleted(ctxTimeout, a.tx, archiveID)
	if err != nil {
		return chunks, err
	}

	for i := 0; i < len(chunks); i++ {
		err = a.attachArchiveToChunk(ctxTimeout, &chunks[i], archiveID)
		if err != nil {
			return chunks, err
		}
	}

	return chunks, nil
}

func (a *archiveUseCase) Store(ctx context.Context, archive *domain.Archive) ([]domain.Chunk, error) {
	if archive.Size > maxArchiveSize {
		return nil, ErrArchiveTooLarge
	}

	if a.tx == nil {
		return nil, ErrNilTx
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	err := a.archiveRepo.Store(ctxTimeout, a.tx, archive)
	if err != nil {
		return nil, fmt.Errorf("storing the archive failed: %w", err)
	}

	chunkNumber := int32(archive.Size / defaultChunkSize)
	lastChunkSize := archive.Size % defaultChunkSize

	currPos := int64(0)
	nextPos := int64(0)

	var chunks []domain.Chunk

	for i := int32(0); i < chunkNumber; i++ {
		currPos = nextPos
		nextPos = currPos + defaultChunkSize

		chunks = append(chunks, domain.Chunk{
			Number:        i,
			StartByte:     currPos,
			NextChunkByte: nextPos,
			Archive:       archive,
		})
	}
	if lastChunkSize != 0 {
		chunks = append(chunks, domain.Chunk{
			Number:        chunkNumber,
			StartByte:     nextPos,
			NextChunkByte: nextPos + lastChunkSize,
			Archive:       archive,
		})
	}

	err = a.chunkRepo.Store(ctxTimeout, a.tx, &chunks)
	if err != nil {
		return nil, fmt.Errorf("chunks storing failed: %w", err)
	}

	return chunks, nil
}

func (a *archiveUseCase) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (a *archiveUseCase) CheckCompleted(ctx context.Context, id int64) (bool, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	return a.archiveRepo.CheckCompleted(ctxTimeout, a.tx, id)
}

func (a *archiveUseCase) SetCompleted(ctx context.Context, id int64, isCompleted bool) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	return a.archiveRepo.SetCompleted(ctxTimeout, a.tx, id, isCompleted)
}

func (a *archiveUseCase) attachArchiveToChunk(ctx context.Context, chunk *domain.Chunk, archiveID int64) error {
	archive, err := a.archiveRepo.GetByID(ctx, a.tx, archiveID)
	if err != nil {
		return fmt.Errorf("load archive failed: %w", err)
	}

	chunk.Archive = &archive
	return nil
}
