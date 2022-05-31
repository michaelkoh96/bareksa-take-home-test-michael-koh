package response

import (
	"bareksa-take-home-test-michael-koh/core/entity"
)

func MapGetNewsResponse(news []entity.News, topics []entity.Topic) GeneralResponse {
	newsResp := make([]News, 0)
	topicMap := make(map[string]string)

	for _, i := range topics {
		topicMap[i.Serial] = i.Title
	}

	for _, n := range news {
		newsResp = append(newsResp, News{
			Serial:      n.Serial,
			Status:      n.Status,
			AuthorName:  n.AuthorName,
			Description: n.Description,
			Title:       n.Title,
			Tags:        n.Tags,
			Topic: Topic{
				Serial: n.TopicSerial,
				Title:  topicMap[n.TopicSerial],
			},
		})
	}

	return GeneralResponse{
		Data:    newsResp,
		Message: "success",
	}
}

func MapGetTagResponse(tags []entity.Tag) GeneralResponse {
	tagsResp := make([]Tag, 0)

	for _, i := range tags {
		tagsResp = append(tagsResp, Tag{
			Name: i.Name,
		})
	}

	return GeneralResponse{
		Data:    tagsResp,
		Message: "success",
	}
}

func CreatedNewsResponse() GeneralResponse {
	return GeneralResponse{
		Data:    nil,
		Message: "success",
	}
}
