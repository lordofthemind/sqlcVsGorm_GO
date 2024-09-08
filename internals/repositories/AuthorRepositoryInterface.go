package repositories

import (
	"context"

	"github.com/lordofthemind/sqlcVsGorm_GO/internals/sqlc/sqlcgen"
)

// AuthorRepository defines the interface for interacting with authors.
type AuthorRepository interface {
	CreateAuthor(ctx context.Context, name string, bio string) (int64, error)
	GetAuthor(ctx context.Context, id int64) (*sqlcgen.Author, error)
	ListAuthors(ctx context.Context) ([]sqlcgen.Author, error)
	DeleteAuthor(ctx context.Context, id int64) error
}
