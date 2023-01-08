package usecase

import (
	"context"
	"time"

	"github.com/kaz-as/zip/domain"
)

type archiveUseCase struct {
	archiveRepo domain.ArchiveRepository
	chunkRepo   domain.ChunkRepository
	ctxTimeout  time.Duration
}

func New(archiveRepo domain.ArchiveRepository, chunkRepo domain.ChunkRepository, timeout time.Duration) domain.ArchiveUseCase {
	return &archiveUseCase{
		archiveRepo: archiveRepo,
		chunkRepo:   chunkRepo,
		ctxTimeout:  timeout,
	}
}

func (a *archiveUseCase) GetByID(ctx context.Context, id int64) (domain.Archive, error) {
	//TODO implement me
	panic("implement me")
}

func (a *archiveUseCase) GetChunkForUpdate(ctx context.Context, archiveID int64, chunkNumber int32) (domain.Chunk, error) {
	//TODO implement me
	panic("implement me")
}

func (a *archiveUseCase) GetUncompletedChunks(ctx context.Context, archiveID int64) ([]domain.Chunk, error) {
	//TODO implement me
	panic("implement me")
}

func (a *archiveUseCase) Store(ctx context.Context, archive *domain.Archive) ([]domain.Chunk, error) {
	//TODO implement me
	panic("implement me")
}

func (a *archiveUseCase) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (a *archiveUseCase) GetFiles(ctx context.Context, id int64) ([]domain.File, error) {
	//TODO implement me
	panic("implement me")
}

func (a *archiveUseCase) GetFile(ctx context.Context, id int64, path []string, name string) (domain.File, error) {
	//TODO implement me
	panic("implement me")
}
