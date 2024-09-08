package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lordofthemind/sqlcVsGorm_GO/internals/sqlc/sqlcgen"
)

// AuthorRepository defines the methods that both GORM and SQLC repositories must implement.
type AuthorRepository interface {
	CreateAuthor(ctx context.Context, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) (uuid.UUID, error)
	GetAuthor(ctx context.Context, id uuid.UUID) (sqlcgen.Author, error)
	ListAuthors(ctx context.Context) ([]sqlcgen.Author, error)
	DeleteAuthor(ctx context.Context, id uuid.UUID) error
	UpdateAuthor(ctx context.Context, id uuid.UUID, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) error
	GetAuthorsByBirthdateRange(ctx context.Context, startDate, endDate time.Time) ([]sqlcgen.Author, error)
}
