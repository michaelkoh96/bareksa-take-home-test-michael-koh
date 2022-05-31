package news

import (
	"bareksa-take-home-test-michael-koh/core/entity"
	"bareksa-take-home-test-michael-koh/core/repository"
	"context"
)

type (
	NewsService interface {
		GetNews(ctx context.Context, newsQuery entity.GetNewsQuery) ([]entity.News, error)
	}

	newsService struct {
		repo repository.NewsRepository
	}
)

func NewService(repo repository.NewsRepository) NewsService {
	return &newsService{
		repo: repo,
	}
}

func (s *newsService) GetNews(ctx context.Context, newsQuery entity.GetNewsQuery) ([]entity.News, error) {
	news, err := s.repo.GetNewsByQuery(newsQuery)
	if err != nil {
		return nil, err
	}

	return news, nil
}
