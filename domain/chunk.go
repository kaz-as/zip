package domain

import (
	"context"
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
	GetByNumber(archiveID int64, number int32) (Chunk, error)
	Update(context.Context, *Chunk) error
	Store(context.Context, *Chunk) error
	DeleteAll(ctx context.Context, archiveID int64) error
}
