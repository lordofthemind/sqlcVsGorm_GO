package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/lordofthemind/sqlcVsGorm_GO/internals/sqlc/sqlcgen"
	"gorm.io/gorm"
)

type GORMRepository struct {
	db *gorm.DB
}

func NewGORMRepository(db *gorm.DB) *GORMRepository {
	return &GORMRepository{db: db}
}

func (r *GORMRepository) CreateAuthor(ctx context.Context, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) (int32, error) {
	author := sqlcgen.Author{
		Name:        name,
		Bio:         bio,
		Email:       email,
		DateOfBirth: dateOfBirth,
	}
	result := r.db.WithContext(ctx).Create(&author)
	return author.ID, result.Error
}

func (r *GORMRepository) GetAuthor(ctx context.Context, id int32) (sqlcgen.Author, error) {
	var author sqlcgen.Author
	result := r.db.WithContext(ctx).First(&author, id)
	return author, result.Error
}

func (r *GORMRepository) ListAuthors(ctx context.Context) ([]sqlcgen.Author, error) {
	var authors []sqlcgen.Author
	result := r.db.WithContext(ctx).Order("name").Find(&authors)
	return authors, result.Error
}

func (r *GORMRepository) DeleteAuthor(ctx context.Context, id int32) error {
	result := r.db.WithContext(ctx).Delete(&sqlcgen.Author{}, id)
	return result.Error
}

func (r *GORMRepository) UpdateAuthor(ctx context.Context, id int32, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) error {
	author := sqlcgen.Author{
		ID:          id,
		Name:        name,
		Bio:         bio,
		Email:       email,
		DateOfBirth: dateOfBirth,
	}
	result := r.db.WithContext(ctx).Model(&author).Updates(author)
	return result.Error
}

func (r *GORMRepository) GetAuthorsByBirthdateRange(ctx context.Context, startDate, endDate time.Time) ([]sqlcgen.Author, error) {
	var authors []sqlcgen.Author
	result := r.db.WithContext(ctx).
		Where("date_of_birth BETWEEN ? AND ?", startDate, endDate).
		Order("date_of_birth").
		Find(&authors)
	return authors, result.Error
}
