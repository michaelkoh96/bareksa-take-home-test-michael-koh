package repository

import (
	"bareksa-take-home-test-michael-koh/core/entity"
	"errors"
)

var (
	ErrTopicsNotFound = errors.New("error, topic not found")
)

type (
	TopicRepository interface {
		GetTopicsBySerials(serials []string) ([]entity.Topic, error)
	}
)
