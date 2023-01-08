package domain

import (
	"context"
	"time"
)

type Archive struct {
	ID          int64
	DirForFiles string // Directory for unarchived files
	Original    string // Full path of the temporary archive (until it is fully unarchived)
	Size        int64
	Name        string
	InitTime    time.Time
}

type File struct {
	RelativePath []string // Relative path to a folder containing the file
	Name         string
	Archive      *Archive
}

type ArchiveRepository interface {
	GetByID(ctx context.Context, id int64) (Archive, error)
	Update(context.Context, *Archive) error
	Store(context.Context, *Archive) error
	Delete(ctx context.Context, id int64) error
}

type ArchiveUseCase interface {
	GetByID(ctx context.Context, id int64) (Archive, error)
	GetChunkForUpdate(ctx context.Context, archiveID int64, chunkNumber int32) (Chunk, error)
	Store(context.Context, *Archive) ([]Chunk, error)
	Delete(ctx context.Context, id int64) error
	GetFiles(ctx context.Context, id int64) ([]File, error)
	GetFile(ctx context.Context, id int64, path []string, name string) (File, error)
}
