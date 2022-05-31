package tag

import (
	"bareksa-take-home-test-michael-koh/core/repository"
	"context"
)

type (
	NewsService interface {
		DeleteTag(ctx context.Context, tagName string) error
	}

	newsService struct {
		repo repository.TagRepository
	}
)

func NewService(repo repository.TagRepository) NewsService {
	return &newsService{
		repo: repo,
	}
}

func (s *newsService) DeleteTag(ctx context.Context, tagName string) error {
	return s.repo.DeleteTag(tagName)
}
