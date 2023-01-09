package archive_repo

import (
	"context"
	"database/sql"

	"github.com/kaz-as/zip/domain"
)

type postgresArchiveRepo struct {
	DB *sql.DB
}

func New(db *sql.DB) domain.ArchiveRepository {
	return postgresArchiveRepo{
		DB: db,
	}
}

func (p postgresArchiveRepo) GetByID(ctx context.Context, tx *sql.Tx, id int64) (domain.Archive, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresArchiveRepo) Update(ctx context.Context, tx *sql.Tx, archive *domain.Archive) error {
	//TODO implement me
	panic("implement me")
}

func (p postgresArchiveRepo) Store(ctx context.Context, tx *sql.Tx, archive *domain.Archive) error {
	//TODO implement me
	panic("implement me")
}

func (p postgresArchiveRepo) Delete(ctx context.Context, tx *sql.Tx, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p postgresArchiveRepo) CheckCompleted(ctx context.Context, tx *sql.Tx, id int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresArchiveRepo) SetCompleted(ctx context.Context, tx *sql.Tx, id int64, isCompleted bool) error {
	//TODO implement me
	panic("implement me")
}
