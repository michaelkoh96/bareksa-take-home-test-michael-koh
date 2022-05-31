package news

import (
	repo "bareksa-take-home-test-michael-koh/core/repository"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.NewsRepository {
	return &repository{db}
}

var (
	StatusDeleted = "deleted"
)

func (r *repository) UpdateNewsRepo(news News, tags []string) error {
	tx := r.db.Begin()

	err := tx.Table(NewsTableName).Where("serial = ?", news.Serial).Updates(map[string]interface{}{
		"topic_serial": news.TopicSerial,
		"status":       news.Status,
		"title":        news.Title,
		"author_name":  news.AuthorName,
		"description":  news.Description,
	}).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repo.ErrNewsNotFound
		}
		return err
	}

	return tx.Commit().Error
}

// Soft Delete
func (r *repository) DeleteNewsRepo(newsSerial string) error {
	tx := r.db.Begin()
	err := tx.Table(NewsTableName).Where("serial = ?", newsSerial).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
		"status":     StatusDeleted,
	}).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repo.ErrNewsNotFound
		}
		return err
	}

	return tx.Commit().Error
}

func (r *repository) CreateNewsRepo(news News, tags []string) error {
	tx := r.db.Begin()

	err := tx.Create(&news).Error
	if err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return err
	}

	for _, t := range tags {
		err = tx.Table(NewsTagTableName).Create(&NewsTag{
			NewsSerial: news.Serial,
			TagName:    t,
		}).Error
		if err != nil {
			tx.Rollback()
			log.Println(err.Error())
			return err
		}
	}

	return tx.Commit().Error
}

func (r *repository) GetTagsByNewsSerialsRepo(newsSerials []string) (map[string][]string, error) {
	newsTags := make([]NewsTag, 0)
	err := r.db.Table(NewsTagTableName).Where("news_serial IN (?)", newsSerials).Find(&newsTags).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repo.ErrNewsNotFound
		}
		return nil, err
	}

	// Map returned rows to map
	newsTagsMap := make(map[string][]string)
	for _, item := range newsTags {
		if _, ok := newsTagsMap[item.NewsSerial]; ok {
			newsTagsMap[item.NewsSerial] = append(newsTagsMap[item.NewsSerial], item.TagName)
		} else {
			newTagArr := make([]string, 0)
			newTagArr = append(newTagArr, item.TagName)
			newsTagsMap[item.NewsSerial] = newTagArr
		}
	}

	return newsTagsMap, nil
}

func (r *repository) GetNewsByQueryRepo(status string, topicSerials []string) ([]News, error) {
	news := make([]News, 0)
	query := r.db.Table(NewsTableName)

	if status != "" {
		query.Where("status = ?", status)
	}

	if len(topicSerials) != 0 {
		query.Where("topic_serial IN (?)", topicSerials)
	}

	query.Where("deleted_at is NULL AND status != 'deleted'")

	err := query.Find(&news).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repo.ErrNewsNotFound
		}
		return nil, err
	}

	return news, nil
}
