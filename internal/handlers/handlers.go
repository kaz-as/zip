package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/kaz-as/zip/domain"
	"github.com/kaz-as/zip/internal/middlewares"
	"github.com/kaz-as/zip/pkg/archive"
	"github.com/kaz-as/zip/pkg/logger"
	"github.com/kaz-as/zip/restapi"
	"github.com/kaz-as/zip/restapi/operations"
)

type handlerSet struct {
	log            logger.Interface
	archiveUseCase domain.ArchiveUseCase
	archiveService archive.Archiver
}

func (s *handlerSet) checkChunksHandler(params operations.CheckChunksParams) middleware.Responder {
	return middleware.NotImplemented("operation CheckChunks has not yet been implemented")
}

func (s *handlerSet) createArchiveHandler(params operations.CreateArchiveParams) middleware.Responder {
	return middleware.NotImplemented("operation CreateArchive has not yet been implemented")
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
	archiveUseCase domain.ArchiveUseCase,
	archiveService archive.Archiver,
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
		log:            l,
		archiveUseCase: archiveUseCase,
		archiveService: archiveService,
	}

	api.CheckChunksHandler = operations.CheckChunksHandlerFunc(hSet.checkChunksHandler)
	api.CreateArchiveHandler = operations.CreateArchiveHandlerFunc(hSet.createArchiveHandler)
	api.GetFilesHandler = operations.GetFilesHandlerFunc(hSet.getFilesHandler)
	api.InitUploadArchiveHandler = operations.InitUploadArchiveHandlerFunc(hSet.initUploadArchiveHandler)
	api.UploadChunkHandler = operations.UploadChunkHandlerFunc(hSet.uploadChunkHandler)

	return api.Serve(middleware.Builder(middlewares.Chain(mws))), nil
}
