package tag_test

import (
	"bareksa-take-home-test-michael-koh/core/entity"
	"bareksa-take-home-test-michael-koh/core/repository/mocks"
	svcTag "bareksa-take-home-test-michael-koh/core/service/tag"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type tagServiceTest struct {
	MockRepo *mocks.TagRepository
	Service  svcTag.TagsService
}

var svcTest tagServiceTest

func init() {
	mockRepo := new(mocks.TagRepository)
	svcTest = tagServiceTest{
		MockRepo: mockRepo,
		Service:  svcTag.NewService(mockRepo),
	}
}

func TestGetNews(t *testing.T) {
	t.Run("get tags success", func(t *testing.T) {
		mockResponse := []entity.Tag{
			{
				Name: "investasi",
			},
			{
				Name: "reksadana",
			},
		}

		mockPage := 2
		mockSize := 4

		svcTest.MockRepo.On("GetNews", mock.Anything, mock.Anything).Return(&mockResponse, nil).Once()
		result, err := svcTest.Service.GetTags(context.Background(), mockPage, mockSize)
		assert.Nil(t, err, "unexpected error")
		assert.NotNil(t, result, "result nil")
		for idx, i := range mockResponse {
			assert.Equal(t, result[idx].Name, i.Name, "Invalid name")
		}
	})

	t.Run("get tags failed", func(t *testing.T) {
		mockResponse := make([]entity.Tag, 0)

		mockPage := 2
		mockSize := 4

		svcTest.MockRepo.On("GetNews", mock.Anything, mock.Anything).Return(&mockResponse, errors.New("get tag err")).Once()
		result, err := svcTest.Service.GetTags(context.Background(), mockPage, mockSize)
		assert.NotNil(t, err, "err should no nill")
		assert.Empty(t, result, "result should be empty")
		for _, i := range mockResponse {
			assert.Equal(t, i.Name, "", "name should be empty")
		}
	})
}

func TestCreateNews(t *testing.T) {
	t.Run("create news success", func(t *testing.T) {
		request := "saham"

		svcTest.MockRepo.On("CreateNews", mock.Anything, mock.Anything).Return(nil).Once()
		err := svcTest.Service.CreateTags(context.Background(), request)
		assert.Nil(t, err, "unexpected err")
	})

	t.Run("create news failed", func(t *testing.T) {
		request := "saham"

		svcTest.MockRepo.On("CreateNews", mock.Anything, mock.Anything).Return(errors.New("create tag error")).Once()
		err := svcTest.Service.CreateTags(context.Background(), request)
		assert.NotNil(t, err, "should return err")
	})
}

func TestUpdateNews(t *testing.T) {
	t.Run("create news success", func(t *testing.T) {
		request := "saham"
		newTagname := "obligasi"

		svcTest.MockRepo.On("UpdateTags", mock.Anything, mock.Anything).Return(nil).Once()
		err := svcTest.Service.UpdateTags(context.Background(), request, newTagname)
		assert.Nil(t, err, "unexpected err")
	})

	t.Run("create news failed", func(t *testing.T) {
		request := "saham"
		newTagname := "obligasi"

		svcTest.MockRepo.On("UpdateTags", mock.Anything, mock.Anything).Return(errors.New("create tag error")).Once()
		err := svcTest.Service.UpdateTags(context.Background(), request, newTagname)
		assert.NotNil(t, err, "should return err")
	})
}
