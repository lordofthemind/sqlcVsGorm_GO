package repositories

import (
	"context"
	"database/sql"

	"github.com/lordofthemind/sqlcVsGorm_GO/internals/sqlc/sqlcgen"
)

// AuthorRepository defines the methods that both GORM and SQLC repositories must implement.
type AuthorRepository interface {
	CreateAuthor(ctx context.Context, name string, bio sql.NullString) (int64, error)
	GetAuthor(ctx context.Context, id int64) (sqlcgen.Author, error)
	ListAuthors(ctx context.Context) ([]sqlcgen.Author, error)
	DeleteAuthor(ctx context.Context, id int64) error
}
