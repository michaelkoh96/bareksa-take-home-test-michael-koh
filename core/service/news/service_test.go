package news_test

import (
	"bareksa-take-home-test-michael-koh/core/entity"
	"bareksa-take-home-test-michael-koh/core/repository/mocks"
	svcNews "bareksa-take-home-test-michael-koh/core/service/news"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	t.Run("get news success", func(t *testing.T) {
		mockResponse := []entity.News{
			{
				Status:      "publish",
				Title:       "Ayo mulai investasi",
				TopicSerial: "TPC-D84H2",
				AuthorName:  "John Doe",
				Description: "story ...",
				Tags:        []string{"investasi", "pemula"},
			},
			{
				Status:      "publish",
				Title:       "Ayo mulai investasi",
				TopicSerial: "TPC-D74U2",
				AuthorName:  "Susan Doe",
				Description: "story ...",
				Tags:        []string{"crypto", "pendidikan"},
			},
			{
				Status:      "publish",
				Title:       "Jangan takut investasi",
				TopicSerial: "TPC-SP40EI",
				AuthorName:  "Peter Doe",
				Description: "story ...",
				Tags:        []string{"investasi", "pengalaman"},
			},
		}
		newsQuery := entity.GetNewsQuery{
			Status:       "publish",
			TopicSerials: []string{"TPC-D84H2", "TPC-D74U2", "TPC-SP40EI"},
		}

		svcTest.MockRepo.On("GetNews", mock.Anything, mock.Anything).Return(&mockResponse, nil).Once()
		result, err := svcTest.Service.GetNews(context.Background(), newsQuery)
		assert.Nil(t, err, "unexpected error")
		assert.NotNil(t, result, "result nil")
		for idx, i := range mockResponse {
			assert.Equal(t, result[idx].TopicSerial, i.TopicSerial, "Invalid topic serial")
			assert.Equal(t, result[idx].Status, i.Status, "Invalid status")
			assert.Equal(t, result[idx].Title, i.Title, "Invalid title")
			assert.Equal(t, result[idx].Description, i.Description, "Invalid description")
			assert.Equal(t, result[idx].AuthorName, i.AuthorName, "Invalid author name")
		}
	})

	t.Run("get news failed, err not nill", func(t *testing.T) {
		mockResponse := make([]entity.News, 0)
		newsQuery := entity.GetNewsQuery{
			Status:       "publish",
			TopicSerials: []string{"TPC-D84H2", "TPC-D74U2", "TPC-SP40EI"},
		}

		svcTest.MockRepo.On("GetNews", mock.Anything, mock.Anything).Return(&mockResponse, errors.New("database error")).Once()
		_, err := svcTest.Service.GetNews(context.Background(), newsQuery)
		assert.NotNil(t, err, "should contain err")
	})
}

func TestCreateNews(t *testing.T) {
	t.Run("create news success", func(t *testing.T) {
		request := entity.News{
			Status:      "publish",
			Title:       "Jangan takut investasi",
			TopicSerial: "TPC-DJ5930",
			AuthorName:  "Ferdy Doe",
			Description: "story ...",
			Tags:        []string{"investasi", "reksadana"},
		}

		svcTest.MockRepo.On("CreateNews", mock.Anything, mock.Anything).Return(nil).Once()
		err := svcTest.Service.CreateNews(context.Background(), request)
		assert.Nil(t, err, "unexpected err")
	})

	t.Run("create news failed", func(t *testing.T) {
		request := entity.News{
			Status:      "publish",
			Title:       "Jangan takut investasi",
			TopicSerial: "TPC-DJ5930",
			AuthorName:  "Ferdy Doe",
			Description: "story ...",
			Tags:        []string{"investasi", "reksadana"},
		}

		svcTest.MockRepo.On("CreateNews", mock.Anything, mock.Anything).Return(errors.New("create error, already exists")).Once()
		err := svcTest.Service.CreateNews(context.Background(), request)
		assert.NotNil(t, err, "should return err")
	})
}

func TestUpdateNews(t *testing.T) {
	t.Run("create news success", func(t *testing.T) {
		request := entity.News{
			Status:      "publish",
			Title:       "Jangan takut investasi",
			TopicSerial: "TPC-DJ5930",
			AuthorName:  "Ferdy Doe",
		}

		svcTest.MockRepo.On("UpdateNews", mock.Anything, mock.Anything).Return(nil).Once()
		err := svcTest.Service.UpdateNews(context.Background(), request)
		assert.Nil(t, err, "unexpected err")
	})

	t.Run("create news failed", func(t *testing.T) {
		request := entity.News{
			Status:      "publish",
			Title:       "Jangan takut investasi",
			TopicSerial: "TPC-DJ5930",
			AuthorName:  "Ferdy Doe",
		}

		svcTest.MockRepo.On("UpdateNews", mock.Anything, mock.Anything).Return(errors.New("create error, already exists")).Once()
		err := svcTest.Service.UpdateNews(context.Background(), request)
		assert.NotNil(t, err, "should return err")
	})
}
