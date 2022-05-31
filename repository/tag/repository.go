package tag

import (
	repo "bareksa-take-home-test-michael-koh/core/repository"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.TagRepository {
	return &repository{db}
}

func (r *repository) DeleteTagRepo(tagName string) error {
	return r.db.Table(TagTableName).Where("name = ?", tagName).Delete(&Tag{}).Error
}
