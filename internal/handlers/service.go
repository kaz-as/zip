package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/kaz-as/zip/domain"
	"github.com/kaz-as/zip/internal/middlewares"
	"github.com/kaz-as/zip/pkg/archive"
	"github.com/kaz-as/zip/pkg/chunkwriter"
	"github.com/kaz-as/zip/pkg/logger"
	"github.com/kaz-as/zip/restapi"
	"github.com/kaz-as/zip/restapi/operations"
)

type HandlerSet struct {
	Log                logger.Interface
	DB                 *sql.DB
	ArchiveUseCaseMake domain.MakeByTx
	ArchiveService     archive.Archiver
	ChunkWriterService *chunkwriter.ChunkService
	FolderForFiles     string
	FolderForArchives  string

	mx          sync.Mutex
	isUnzipping map[int64]struct{}
}

// todo move HandlerSet to an independent package
func NewHandlerSet(
	log logger.Interface,
	db *sql.DB,
	archiveUseCaseMake domain.MakeByTx,
	archiveService archive.Archiver,
	chunkWriterService *chunkwriter.ChunkService,
	folderForFiles string,
	folderForArchives string,
) *HandlerSet {
	return &HandlerSet{
		Log:                log,
		DB:                 db,
		ArchiveUseCaseMake: archiveUseCaseMake,
		ArchiveService:     archiveService,
		ChunkWriterService: chunkWriterService,
		FolderForFiles:     folderForFiles,
		FolderForArchives:  folderForArchives,
		mx:                 sync.Mutex{},
		isUnzipping:        make(map[int64]struct{}),
	}
}

// New creates handler using middlewares running after routing
func New(
	l logger.Interface,
	db *sql.DB,
	archiveUseCaseMake domain.MakeByTx,
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

	hSet := NewHandlerSet(
		l,
		db,
		archiveUseCaseMake,
		archiveService,
		chunkWriterService,
		folderForFiles,
		folderForArchives,
	)

	api.CheckChunksHandler = operations.CheckChunksHandlerFunc(hSet.CheckChunksHandler)
	api.CreateArchiveHandler = operations.CreateArchiveHandlerFunc(hSet.CreateArchiveHandler)
	api.GetFilesHandler = operations.GetFilesHandlerFunc(hSet.GetFilesHandler)
	api.InitUploadArchiveHandler = operations.InitUploadArchiveHandlerFunc(hSet.InitUploadArchiveHandler)
	api.UploadChunkHandler = operations.UploadChunkHandlerFunc(hSet.UploadChunkHandler)
	api.IsCompletedHandler = operations.IsCompletedHandlerFunc(hSet.IsCompletedHandler)

	return api.Serve(middleware.Builder(middlewares.Chain(mws))), nil
}
