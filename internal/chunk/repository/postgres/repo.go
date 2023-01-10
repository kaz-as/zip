package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/kaz-as/zip/domain"
	"github.com/kaz-as/zip/pkg/logger"
)

type postgresChunkRepo struct {
	log logger.Interface
	DB  *sql.DB
}

func New(db *sql.DB, log logger.Interface) domain.ChunkRepository {
	return postgresChunkRepo{
		DB:  db,
		log: log,
	}
}

func (p postgresChunkRepo) GetByNumber(ctx context.Context, tx *sql.Tx, archiveID int64, number int32) (domain.Chunk, error) {
	query := `SELECT * FROM chunks WHERE archive = ? AND number = ?`
	chunks, err := p.fetch(ctx, tx, query, archiveID, number)
	if err != nil {
		return domain.Chunk{}, fmt.Errorf("fetching chunk: %w", err)
	}

	if len(chunks) == 0 {
		return domain.Chunk{}, domain.ErrNotFound
	}

	return chunks[0], nil
}

func (p postgresChunkRepo) GetUncompleted(ctx context.Context, tx *sql.Tx, archiveID int64) ([]domain.Chunk, error) {
	query := `SELECT * FROM chunks WHERE archive = ? AND NOT uploaded`
	return p.fetch(ctx, tx, query, archiveID)
}

func (p postgresChunkRepo) Update(ctx context.Context, tx *sql.Tx, chunk *domain.Chunk) error {
	query := `UPDATE chunks SET start_byte = ?, next_chunk_byte = ?, uploaded = ?, uploaded_time = ? WHERE archive = ? AND number = ?`

	prepareContext := p.DB.PrepareContext
	if tx != nil {
		prepareContext = tx.PrepareContext
	}

	stmt, err := prepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("prepare context: %w", err)
	}

	res, err := stmt.ExecContext(ctx,
		chunk.StartByte, chunk.NextChunkByte, chunk.Uploaded, chunk.UploadedTime, chunk.Archive.ID, chunk.Number)

	if err != nil {
		return fmt.Errorf("exec context: %w", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if affect != 1 {
		return fmt.Errorf("weird behavior: total affected: %d", affect)
	}

	return nil
}

func (p postgresChunkRepo) Store(ctx context.Context, tx *sql.Tx, chunks *[]domain.Chunk) error {
	if chunks == nil || len(*chunks) == 0 {
		return nil
	}

	query := `INSERT INTO chunks (archive, number, start_byte, next_chunk_byte, uploaded, uploaded_time) VALUES `
	var b strings.Builder
	vls := `(?,?,?,?,?)`
	n := (len(vls)+1)*len(*chunks) - 1
	b.Grow(n)
	b.WriteString(vls)
	b.WriteByte(',')
	for b.Len() < n {
		if b.Len() <= n/2 {
			b.WriteString(b.String())
		} else {
			b.WriteString(b.String()[:n-b.Len()])
			break
		}
	}

	query += b.String()

	prepareContext := p.DB.PrepareContext
	if tx != nil {
		prepareContext = tx.PrepareContext
	}

	stmt, err := prepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("prepare context: %w", err)
	}

	var args []any
	for i := 0; i < len(*chunks); i++ {
		args = append(args,
			(*chunks)[i].Archive.ID,
			(*chunks)[i].Number,
			(*chunks)[i].StartByte,
			(*chunks)[i].NextChunkByte,
			(*chunks)[i].Uploaded,
			(*chunks)[i].UploadedTime,
		)
	}

	_, err = stmt.ExecContext(ctx, args...)

	if err != nil {
		return fmt.Errorf("exec context: %w", err)
	}

	return nil
}

func (p postgresChunkRepo) DeleteAll(ctx context.Context, tx *sql.Tx, archiveID int64) error {
	//TODO implement me
	panic("implement me")
}

func (p postgresChunkRepo) fetch(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (result []domain.Chunk, err error) {
	queryContext := p.DB.QueryContext
	if tx != nil {
		queryContext = tx.QueryContext
	}

	rows, err := queryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			p.log.Error(errRow)
		}
	}()

	result = make([]domain.Chunk, 0)
	for rows.Next() {
		t := domain.Chunk{Archive: new(domain.Archive)}
		err = rows.Scan(
			&t.Archive.ID,
			&t.Number,
			&t.StartByte,
			&t.NextChunkByte,
			&t.Uploaded,
			&t.UploadedTime,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}
