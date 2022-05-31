package tag

import (
	repo "bareksa-take-home-test-michael-koh/core/repository"
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.TagRepository {
	return &repository{db}
}

func (r *repository) GetTagsRepo(page, size int) ([]Tag, error) {
	tags := make([]Tag, 0)
	query := r.db.Table(TagTableName).Limit(size).Offset((page - 1) * size)

	err := query.Find(&tags).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repo.ErrTagNotFound
		}
		return nil, err
	}

	return tags, nil
}

func (r *repository) DeleteTagRepo(tagName string) error {
	return r.db.Table(TagTableName).Where("name = ?", tagName).Delete(&Tag{}).Error
}
