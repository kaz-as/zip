package postgres

import (
	"context"
	"database/sql"

	"github.com/kaz-as/zip/domain"
)

type postgresChunkRepo struct {
	DB *sql.DB
}

func New(db *sql.DB) domain.ChunkRepository {
	return postgresChunkRepo{
		DB: db,
	}
}

func (p postgresChunkRepo) GetByNumber(archiveID int64, number int32) (domain.Chunk, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresChunkRepo) GetUncompleted(archiveID int64) ([]domain.Chunk, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresChunkRepo) Update(ctx context.Context, chunk *domain.Chunk) error {
	//TODO implement me
	panic("implement me")
}

func (p postgresChunkRepo) Store(ctx context.Context, chunk *domain.Chunk) error {
	//TODO implement me
	panic("implement me")
}

func (p postgresChunkRepo) DeleteAll(ctx context.Context, archiveID int64) error {
	//TODO implement me
	panic("implement me")
}
