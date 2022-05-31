package repository

import (
	"bareksa-take-home-test-michael-koh/core/entity"
	"errors"
)

var (
	ErrNewsNotFound           = errors.New("error, news not found")
	ErrCreateNewsInvalidTopic = errors.New("error while creating news, invalid topic")
	ErrCreateNewsInvalidTag   = errors.New("error while creating news, invalid tag")
)

type (
	NewsRepository interface {
		GetNewsByQuery(query entity.GetNewsQuery) ([]entity.News, error)
		CreateNews(news entity.News) error
		UpdateNews(news entity.News) error
	}
)
