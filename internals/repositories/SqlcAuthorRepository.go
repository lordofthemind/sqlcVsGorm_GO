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

func (r *SqlcAuthorRepository) CreateAuthor(ctx context.Context, name string, bio string) (int64, error) {
	params := sqlcgen.CreateAuthorParams{
		Name: name,
		Bio:  sql.NullString{String: bio, Valid: bio != ""},
	}

	// Execute the query to insert the author
	result, err := r.queries.CreateAuthor(ctx, params)
	if err != nil {
		return 0, err
	}

	// Extract the last inserted ID from the result
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SqlcAuthorRepository) GetAuthor(ctx context.Context, id int64) (*sqlcgen.Author, error) {
	author, err := r.queries.GetAuthor(ctx, id)
	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *SqlcAuthorRepository) ListAuthors(ctx context.Context) ([]sqlcgen.Author, error) {
	return r.queries.ListAuthors(ctx)
}

func (r *SqlcAuthorRepository) DeleteAuthor(ctx context.Context, id int64) error {
	return r.queries.DeleteAuthor(ctx, id)
}
