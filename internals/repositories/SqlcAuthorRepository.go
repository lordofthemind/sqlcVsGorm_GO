package repositories

import (
	"context"
	"database/sql"

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

func (r *SqlcAuthorRepository) CreateAuthor(ctx context.Context, name string, bio sql.NullString) (int64, error) {
	params := sqlcgen.CreateAuthorParams{
		Name: name,
		Bio:  bio,
	}
	result, err := r.queries.CreateAuthor(ctx, params)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId() // Return the inserted ID
}

func (r *SqlcAuthorRepository) GetAuthor(ctx context.Context, id int64) (sqlcgen.Author, error) {
	return r.queries.GetAuthor(ctx, id)
}

func (r *SqlcAuthorRepository) ListAuthors(ctx context.Context) ([]sqlcgen.Author, error) {
	return r.queries.ListAuthors(ctx)
}

func (r *SqlcAuthorRepository) DeleteAuthor(ctx context.Context, id int64) error {
	return r.queries.DeleteAuthor(ctx, id)
}
