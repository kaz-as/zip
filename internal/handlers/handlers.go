package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/kaz-as/zip/domain"
	"github.com/kaz-as/zip/internal/middlewares"
	"github.com/kaz-as/zip/models"
	"github.com/kaz-as/zip/pkg/archive"
	"github.com/kaz-as/zip/pkg/chunkwriter"
	"github.com/kaz-as/zip/pkg/logger"
	"github.com/kaz-as/zip/restapi"
	"github.com/kaz-as/zip/restapi/operations"
)

type handlerSet struct {
	log                logger.Interface
	db                 *sql.DB
	archiveUseCase     domain.ArchiveUseCase
	archiveService     archive.Archiver
	chunkWriterService *chunkwriter.ChunkService
	folderForFiles     string
	folderForArchives  string
}

func (s *handlerSet) checkChunksHandler(params operations.CheckChunksParams) middleware.Responder {
	chunks, err := s.archiveUseCase.GetUncompletedChunks(params.HTTPRequest.Context(), params.ID)
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

func (s *handlerSet) createArchiveHandler(params operations.CreateArchiveParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	var filesForArchive []*archive.File

	retError := func() middleware.Responder {
		for _, fp := range filesForArchive {
			if fp == nil {
				continue
			}
			err := fp.In.Close()
			if err != nil {
				s.log.Warn("file close: %s", err)
			}
		}
		return operations.NewCreateArchiveDefault(0)
	}

	for _, a := range params.Files {
		inputFiles := a.Files
		archiveID := a.ID

		currArchive, err := s.archiveUseCase.GetByID(ctx, int64(archiveID))
		if err != nil {
			s.log.Info("s.archiveUseCase.GetByID: %s", err)
			return retError()
		}

		for _, file := range inputFiles {
			path := []string{s.folderForFiles, currArchive.DirForFiles}
			path = append(path, file.Path...)
			path = append(path, string(file.Name))

			finalFilePath := filepath.Join(path...)
			f, err := os.Open(finalFilePath)
			if err != nil {
				s.log.Info("os.Open(\"%s\"): %s", finalFilePath, err)
				return retError()
			}

			filesForArchive = append(filesForArchive, &archive.File{
				In:          f,
				OutLocation: file.NewPath,
				OutName:     string(file.NewName),
			})
		}
	}

	zipOut, err := s.archiveService.Zip(ctx, filesForArchive...)
	if err != nil {
		s.log.Info("zip: %s", err)
		return retError()
	}

	return operations.NewCreateArchiveOK().WithPayload(zipOut)
}

func (s *handlerSet) getFilesHandler(params operations.GetFilesParams) middleware.Responder {
	return middleware.NotImplemented("operation GetFiles has not yet been implemented")
}

func (s *handlerSet) initUploadArchiveHandler(params operations.InitUploadArchiveParams) middleware.Responder {
	return middleware.NotImplemented("operation InitUploadArchive has not yet been implemented")
}

func (s *handlerSet) uploadChunkHandler(params operations.UploadChunkParams) middleware.Responder {
	return middleware.NotImplemented("operation UploadChunk has not yet been implemented")
}

// New creates handler using middlewares running after routing
func New(
	l logger.Interface,
	db *sql.DB,
	archiveUseCase domain.ArchiveUseCase,
	archiveService archive.Archiver,
	chunkWriterService *chunkwriter.ChunkService,
	folderForFiles string,
	folderForArchives string,
	mws []middlewares.Middleware,
) (http.Handler, error) {
	swaggerDoc, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, fmt.Errorf("swagger loading: %s", err)
	}

	api := operations.NewZipAPI(swaggerDoc)
	api.Logger = l.Info
	api.UseSwaggerUI()

	hSet := handlerSet{
		log:                l,
		db:                 db,
		archiveUseCase:     archiveUseCase,
		archiveService:     archiveService,
		chunkWriterService: chunkWriterService,
		folderForFiles:     folderForFiles,
		folderForArchives:  folderForArchives,
	}

	api.CheckChunksHandler = operations.CheckChunksHandlerFunc(hSet.checkChunksHandler)
	api.CreateArchiveHandler = operations.CreateArchiveHandlerFunc(hSet.createArchiveHandler)
	api.GetFilesHandler = operations.GetFilesHandlerFunc(hSet.getFilesHandler)
	api.InitUploadArchiveHandler = operations.InitUploadArchiveHandlerFunc(hSet.initUploadArchiveHandler)
	api.UploadChunkHandler = operations.UploadChunkHandlerFunc(hSet.uploadChunkHandler)

	return api.Serve(middleware.Builder(middlewares.Chain(mws))), nil
}
