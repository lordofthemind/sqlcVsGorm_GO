package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lordofthemind/sqlcVsGorm_GO/internals/sqlc/sqlcgen"
)

type SqlcAuthorRepository struct {
	queries *sqlcgen.Queries
}

func NewSqlcAuthorRepository(db *sql.DB) *SqlcAuthorRepository {
	return &SqlcAuthorRepository{
		queries: sqlcgen.New(db),
	}
}

func (r *SqlcAuthorRepository) CreateAuthor(ctx context.Context, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) (uuid.UUID, error) {
	params := sqlcgen.CreateAuthorParams{
		Name:        name,
		Bio:         bio,
		Email:       email,
		DateOfBirth: dateOfBirth,
	}
	return r.queries.CreateAuthor(ctx, params)
}

func (r *SqlcAuthorRepository) GetAuthor(ctx context.Context, id uuid.UUID) (sqlcgen.Author, error) {
	return r.queries.GetAuthor(ctx, id)
}

func (r *SqlcAuthorRepository) ListAuthors(ctx context.Context) ([]sqlcgen.Author, error) {
	return r.queries.ListAuthors(ctx)
}

func (r *SqlcAuthorRepository) DeleteAuthor(ctx context.Context, id uuid.UUID) error {
	return r.queries.DeleteAuthor(ctx, id)
}

func (r *SqlcAuthorRepository) UpdateAuthor(ctx context.Context, id uuid.UUID, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) error {
	params := sqlcgen.UpdateAuthorParams{
		ID:          id,
		Name:        name,
		Bio:         bio,
		Email:       email,
		DateOfBirth: dateOfBirth,
	}
	return r.queries.UpdateAuthor(ctx, params)
}

func (r *SqlcAuthorRepository) GetAuthorsByBirthdateRange(ctx context.Context, startDate, endDate time.Time) ([]sqlcgen.Author, error) {
	params := sqlcgen.GetAuthorsByBirthdateRangeParams{
		DateOfBirth:   sql.NullTime{Time: startDate, Valid: true},
		DateOfBirth_2: sql.NullTime{Time: endDate, Valid: true},
	}
	return r.queries.GetAuthorsByBirthdateRange(ctx, params)
}
