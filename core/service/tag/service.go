package tag

import (
	"bareksa-take-home-test-michael-koh/core/entity"
	"bareksa-take-home-test-michael-koh/core/repository"
	"context"
)

type (
	TagsService interface {
		GetTags(ctx context.Context, page, size int) ([]entity.Tag, error)
		DeleteTag(ctx context.Context, tagName string) error
	}

	tagService struct {
		repo repository.TagRepository
	}
)

func NewService(repo repository.TagRepository) TagsService {
	return &tagService{
		repo: repo,
	}
}

func (s *tagService) DeleteTag(ctx context.Context, tagName string) error {
	return s.repo.DeleteTag(tagName)
}

func (s *tagService) GetTags(ctx context.Context, page, size int) ([]entity.Tag, error) {
	return s.repo.GetTags(page, size)
}
