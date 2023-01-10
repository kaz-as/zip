package domain

import (
	"context"
	"database/sql"
	"time"
)

type Archive struct {
	ID int64
	// A name of the archive on the hard drive (until it is fully unarchived) and of the folder with unarchived files
	Original   string
	Size       int64
	Name       string
	InitTime   time.Time
	Unarchived bool
}

type ArchiveRepository interface {
	GetByID(ctx context.Context, tx *sql.Tx, id int64) (Archive, error)
	Update(context.Context, *sql.Tx, *Archive) error
	Store(context.Context, *sql.Tx, *Archive) error
	Delete(ctx context.Context, tx *sql.Tx, id int64) error
	CheckCompleted(ctx context.Context, tx *sql.Tx, id int64) (bool, error)
	SetCompleted(ctx context.Context, tx *sql.Tx, id int64, isCompleted bool) error
}

type ArchiveUseCase interface {
	GetByID(ctx context.Context, id int64) (Archive, error)
	GetChunkForUpdate(ctx context.Context, archiveID int64, chunkNumber int32) (Chunk, error)
	UpdateChunk(context.Context, *Chunk) error
	GetUncompletedChunks(ctx context.Context, archiveID int64) ([]Chunk, error)
	Store(context.Context, *Archive) ([]Chunk, error)
	Delete(ctx context.Context, id int64) error
	CheckCompleted(ctx context.Context, id int64) (bool, error)
	SetCompleted(ctx context.Context, id int64, isCompleted bool) error
}

type MakeByTx func(*sql.Tx) ArchiveUseCase
