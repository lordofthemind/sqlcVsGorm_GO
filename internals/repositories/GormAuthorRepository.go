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

func (r *GormAuthorRepository) CreateAuthor(ctx context.Context, name string, bio string) (int64, error) {
	author := sqlcgen.Author{
		Name: name,
		Bio:  sql.NullString{String: bio, Valid: bio != ""}, // Correctly setting the Bio field as sql.NullString
	}

	if err := r.db.WithContext(ctx).Create(&author).Error; err != nil {
		return 0, err
	}

	return author.ID, nil
}

func (r *GormAuthorRepository) GetAuthor(ctx context.Context, id int64) (*sqlcgen.Author, error) {
	var author sqlcgen.Author
	if err := r.db.WithContext(ctx).First(&author, id).Error; err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *GormAuthorRepository) ListAuthors(ctx context.Context) ([]sqlcgen.Author, error) {
	var authors []sqlcgen.Author
	if err := r.db.WithContext(ctx).Order("name").Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *GormAuthorRepository) DeleteAuthor(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&sqlcgen.Author{}, id).Error
}
