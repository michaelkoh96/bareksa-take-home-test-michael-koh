package repository

import (
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
	}
)
