package domain

import (
	"context"
	"database/sql"
	"time"
)

type Chunk struct {
	Number        int32
	StartByte     int64
	NextChunkByte int64
	Uploaded      bool
	UploadedTime  time.Time
	Archive       *Archive
}

type ChunkRepository interface {
	GetByNumber(ctx context.Context, tx *sql.Tx, archiveID int64, number int32) (Chunk, error)
	GetUncompleted(ctx context.Context, tx *sql.Tx, archiveID int64) ([]Chunk, error)
	Update(context.Context, *sql.Tx, *Chunk) error
	Store(context.Context, *sql.Tx, *Chunk) error
	DeleteAll(ctx context.Context, tx *sql.Tx, archiveID int64) error
}
