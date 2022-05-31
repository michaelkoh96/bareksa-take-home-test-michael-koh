package repository

import (
	"bareksa-take-home-test-michael-koh/core/entity"
	"errors"
)

var (
	ErrTagNotFound = errors.New("error, news not found")
)

type (
	TagRepository interface {
		// GetNewsByQuery(query entity.GetNewsQuery) ([]entity.News, error)
		// CreateNews(news entity.News) error
		// UpdateNews(news entity.News) error
		DeleteTag(tagName string) error
		GetTags(page, size int) ([]entity.Tag, error)
	}
)
