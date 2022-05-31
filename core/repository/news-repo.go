package repository

import (
	"bareksa-take-home-test-michael-koh/core/entity"
	"errors"
)

var (
	ErrNewsNotFound = errors.New("error, news not found")
)

type (
	NewsRepository interface {
		GetNewsByQuery(query entity.GetNewsQuery) ([]entity.News, error)
	}
)
