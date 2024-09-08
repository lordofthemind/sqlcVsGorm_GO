package repositories

import (
	"context"
	"database/sql"

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

func (r *GormAuthorRepository) CreateAuthor(ctx context.Context, name string, bio sql.NullString) (int64, error) {
	author := sqlcgen.Author{
		Name: name,
		Bio:  bio,
	}
	result := r.db.WithContext(ctx).Create(&author)
	if result.Error != nil {
		return 0, result.Error
	}
	return author.ID, nil // Return the inserted ID
}

func (r *GormAuthorRepository) GetAuthor(ctx context.Context, id int64) (sqlcgen.Author, error) {
	var author sqlcgen.Author
	err := r.db.WithContext(ctx).First(&author, id).Error
	return author, err
}

func (r *GormAuthorRepository) ListAuthors(ctx context.Context) ([]sqlcgen.Author, error) {
	var authors []sqlcgen.Author
	err := r.db.WithContext(ctx).Find(&authors).Error
	return authors, err
}

func (r *GormAuthorRepository) DeleteAuthor(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&sqlcgen.Author{}, id).Error
}
