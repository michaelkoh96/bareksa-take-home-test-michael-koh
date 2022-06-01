package news_test

import (
	"bareksa-take-home-test-michael-koh/core/repository/mocks"
	svcNews "bareksa-take-home-test-michael-koh/core/service/news"
	"testing"
)

type newsServiceTest struct {
	MockRepo *mocks.NewsRepository
	Service  svcNews.NewsService
}

var svcTest newsServiceTest

func init() {
	mockRepo := new(mocks.NewsRepository)
	svcTest = newsServiceTest{
		MockRepo: mockRepo,
		Service:  svcNews.NewService(mockRepo),
	}
}

func TestGetNews(t *testing.T) {

}
