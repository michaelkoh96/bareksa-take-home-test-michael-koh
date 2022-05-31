package topic

import (
	repo "bareksa-take-home-test-michael-koh/core/repository"
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.TopicRepository {
	return &repository{db}
}

func (r *repository) GetTopicsBySerialsRepo(serials []string) ([]Topic, error) {
	topics := make([]Topic, 0)
	err := r.db.Table(TopicTableName).Where("serial IN (?)", serials).Find(&topics).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repo.ErrTopicsNotFound
		}
		return nil, err
	}

	return topics, nil
}
