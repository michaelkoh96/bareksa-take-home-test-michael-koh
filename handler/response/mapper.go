package response

import (
	"bareksa-take-home-test-michael-koh/core/entity"
)

func MapGetNewsResponse(news []entity.News) GeneralResponse {
	newsResp := make([]News, 0)
	for _, n := range news {
		newsResp = append(newsResp, News{
			Serial:      n.Serial,
			Status:      n.Status,
			AuthorName:  n.AuthorName,
			Description: n.Description,
			Title:       n.Title,
			Tags:        n.Tags,
		})
	}

	return GeneralResponse{
		Data:    newsResp,
		Message: "success",
	}
}

func MapErrorResponse(err error) GeneralResponse {
	return GeneralResponse{
		Data:    nil,
		Message: err.Error(),
	}
}
