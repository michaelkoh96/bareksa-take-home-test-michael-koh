package topic

import (
	"bareksa-take-home-test-michael-koh/core/entity"
	"bareksa-take-home-test-michael-koh/core/repository"
	"context"
)

type (
	TopicService interface {
		GetTopicsBySerials(ctx context.Context, serials []string) ([]entity.Topic, error)
	}

	topicService struct {
		repo repository.TopicRepository
	}
)

func NewService(repo repository.TopicRepository) TopicService {
	return &topicService{
		repo: repo,
	}
}

func (s *topicService) GetTopicsBySerials(ctx context.Context, serials []string) ([]entity.Topic, error) {
	topics, err := s.repo.GetTopicsBySerials(serials)
	if err != nil {
		return nil, err
	}

	return topics, nil
}
