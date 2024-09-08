package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/lordofthemind/sqlcVsGorm_GO/internals/sqlc/sqlcgen"
)

// AuthorRepository defines the methods that both GORM and SQLC repositories must implement.
type AuthorRepository interface {
	CreateAuthor(ctx context.Context, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) (int32, error)
	GetAuthor(ctx context.Context, id int32) (sqlcgen.Author, error)
	ListAuthors(ctx context.Context) ([]sqlcgen.Author, error)
	DeleteAuthor(ctx context.Context, id int32) error
	UpdateAuthor(ctx context.Context, id int32, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) error
	GetAuthorsByBirthdateRange(ctx context.Context, startDate, endDate time.Time) ([]sqlcgen.Author, error)
}
