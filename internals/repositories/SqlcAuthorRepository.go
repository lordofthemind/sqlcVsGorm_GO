package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/lordofthemind/sqlcVsGorm_GO/internals/sqlc/sqlcgen"
)

type SQLCRepository struct {
	queries *sqlcgen.Queries
}

func NewSQLCRepository(queries *sqlcgen.Queries) *SQLCRepository {
	return &SQLCRepository{queries: queries}
}

func (r *SQLCRepository) CreateAuthor(ctx context.Context, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) (int32, error) {
	params := sqlcgen.CreateAuthorParams{
		Name:        name,
		Bio:         bio,
		Email:       email,
		DateOfBirth: dateOfBirth,
	}
	return r.queries.CreateAuthor(ctx, params)
}

func (r *SQLCRepository) GetAuthor(ctx context.Context, id int32) (sqlcgen.Author, error) {
	return r.queries.GetAuthor(ctx, id)
}

func (r *SQLCRepository) ListAuthors(ctx context.Context) ([]sqlcgen.Author, error) {
	return r.queries.ListAuthors(ctx)
}

func (r *SQLCRepository) DeleteAuthor(ctx context.Context, id int32) error {
	return r.queries.DeleteAuthor(ctx, id)
}

func (r *SQLCRepository) UpdateAuthor(ctx context.Context, id int32, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) error {
	params := sqlcgen.UpdateAuthorParams{
		ID:          id,
		Name:        name,
		Bio:         bio,
		Email:       email,
		DateOfBirth: dateOfBirth,
	}
	return r.queries.UpdateAuthor(ctx, params)
}

func (r *SQLCRepository) GetAuthorsByBirthdateRange(ctx context.Context, startDate, endDate time.Time) ([]sqlcgen.Author, error) {
	params := sqlcgen.GetAuthorsByBirthdateRangeParams{
		DateOfBirth:   sql.NullTime{Time: startDate, Valid: true},
		DateOfBirth_2: sql.NullTime{Time: endDate, Valid: true},
	}
	return r.queries.GetAuthorsByBirthdateRange(ctx, params)
}
