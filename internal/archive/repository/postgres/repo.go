package archive_repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kaz-as/zip/domain"
	"github.com/kaz-as/zip/pkg/logger"
)

type postgresArchiveRepo struct {
	log logger.Interface
	DB  *sql.DB
}

func New(db *sql.DB, log logger.Interface) domain.ArchiveRepository {
	return postgresArchiveRepo{
		DB:  db,
		log: log,
	}
}

func (p postgresArchiveRepo) GetByID(ctx context.Context, tx *sql.Tx, id int64) (domain.Archive, error) {
	query := `SELECT * FROM archives WHERE id = ?`
	archives, err := p.fetch(ctx, tx, query, id)
	if err != nil {
		return domain.Archive{}, fmt.Errorf("fetching archive: %w", err)
	}

	if len(archives) == 0 {
		return domain.Archive{}, domain.ErrNotFound
	}

	return archives[0], nil
}

func (p postgresArchiveRepo) Update(ctx context.Context, tx *sql.Tx, archive *domain.Archive) error {
	query := `UPDATE archives SET size = ?, name = ?, original = ?, unarchived = ?, created = ? WHERE id = ?`

	prepareContext := p.DB.PrepareContext
	if tx != nil {
		prepareContext = tx.PrepareContext
	}

	stmt, err := prepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("prepare context: %w", err)
	}

	res, err := stmt.ExecContext(ctx,
		archive.Size, archive.Name, archive.Original, archive.Unarchived, archive.InitTime, archive.ID)

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

func (p postgresArchiveRepo) Store(ctx context.Context, tx *sql.Tx, archive *domain.Archive) error {
	query := `INSERT INTO archives (size, name, original, unarchived, created) VALUES (?, ?, ?, ?, ?)`

	prepareContext := p.DB.PrepareContext
	if tx != nil {
		prepareContext = tx.PrepareContext
	}

	stmt, err := prepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("prepare context: %w", err)
	}

	res, err := stmt.ExecContext(ctx,
		archive.Size, archive.Name, archive.Original, archive.Unarchived, archive.InitTime)

	if err != nil {
		return fmt.Errorf("exec context: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("last insert id: %w", err)
	}

	archive.ID = id

	return nil
}

func (p postgresArchiveRepo) Delete(ctx context.Context, tx *sql.Tx, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p postgresArchiveRepo) CheckCompleted(ctx context.Context, tx *sql.Tx, id int64) (res bool, err error) {
	queryContext := p.DB.QueryContext
	if tx != nil {
		queryContext = tx.QueryContext
	}

	query := `SELECT unarchived FROM archives WHERE id = ?`

	rows, err := queryContext(ctx, query, id)
	if err != nil {
		return false, fmt.Errorf("query context: %w", err)
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			p.log.Error(errRow)
		}
	}()

	if !rows.Next() {
		return false, domain.ErrNotFound
	}

	err = rows.Scan(&res)
	if err != nil {
		return false, fmt.Errorf("scan: %w", err)
	}

	if rows.Next() {
		return false, fmt.Errorf("wierd behaviour: more than one row selected")
	}

	return
}

func (p postgresArchiveRepo) SetCompleted(ctx context.Context, tx *sql.Tx, id int64, isCompleted bool) error {
	query := `UPDATE archives SET unarchived = ? WHERE id = ?`

	prepareContext := p.DB.PrepareContext
	if tx != nil {
		prepareContext = tx.PrepareContext
	}

	stmt, err := prepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("prepare context: %w", err)
	}

	res, err := stmt.ExecContext(ctx,
		isCompleted, id)

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

func (p postgresArchiveRepo) fetch(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (result []domain.Archive, err error) {
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

	result = make([]domain.Archive, 0)
	for rows.Next() {
		t := domain.Archive{}
		err = rows.Scan(
			&t.ID,
			&t.Size,
			&t.Name,
			&t.Original,
			&t.Unarchived,
			&t.InitTime,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}
