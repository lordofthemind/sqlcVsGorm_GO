package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lordofthemind/sqlcVsGorm_GO/internals/sqlc/sqlcgen"
	"gorm.io/gorm"
)

type GormAuthorRepository struct {
	db *gorm.DB
}

func NewGormAuthorRepository(db *gorm.DB) *GormAuthorRepository {
	return &GormAuthorRepository{
		db: db,
	}
}

func (r *GormAuthorRepository) CreateAuthor(ctx context.Context, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) (uuid.UUID, error) {
	author := sqlcgen.Author{
		Name:        name,
		Bio:         bio,
		Email:       email,
		DateOfBirth: dateOfBirth,
	}
	result := r.db.WithContext(ctx).Create(&author)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return author.ID, nil
}

func (r *GormAuthorRepository) GetAuthor(ctx context.Context, id uuid.UUID) (sqlcgen.Author, error) {
	var author sqlcgen.Author
	err := r.db.WithContext(ctx).First(&author, "id = ?", id).Error
	return author, err
}

func (r *GormAuthorRepository) ListAuthors(ctx context.Context) ([]sqlcgen.Author, error) {
	var authors []sqlcgen.Author
	err := r.db.WithContext(ctx).Find(&authors).Error
	return authors, err
}

func (r *GormAuthorRepository) DeleteAuthor(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&sqlcgen.Author{}, "id = ?", id).Error
}

func (r *GormAuthorRepository) UpdateAuthor(ctx context.Context, id uuid.UUID, name string, bio sql.NullString, email string, dateOfBirth sql.NullTime) error {
	return r.db.WithContext(ctx).Model(&sqlcgen.Author{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":          name,
		"bio":           bio,
		"email":         email,
		"date_of_birth": dateOfBirth,
	}).Error
}

func (r *GormAuthorRepository) GetAuthorsByBirthdateRange(ctx context.Context, startDate, endDate time.Time) ([]sqlcgen.Author, error) {
	var authors []sqlcgen.Author
	err := r.db.WithContext(ctx).Where("date_of_birth BETWEEN ? AND ?", startDate, endDate).Order("date_of_birth").Find(&authors).Error
	return authors, err
}
