package news

import (
	"bareksa-take-home-test-michael-koh/core/entity"
)

func (r *repository) GetNewsByQuery(query entity.GetNewsQuery) ([]entity.News, error) {
	news, err := r.GetNewsByQueryRepo(query.Status, query.TopicSerials)
	if err != nil {
		return nil, err
	}

	newsSerial := make([]string, 0)
	for _, n := range news {
		newsSerial = append(newsSerial, n.Serial)
	}

	tagMaps, err := r.GetTagsByNewsSerialsRepo(newsSerial)
	if err != nil {
		return nil, err
	}

	newsRes := make([]entity.News, 0)
	for _, item := range news {
		newsRes = append(newsRes, entity.News{
			Serial:      item.Serial,
			TopicSerial: item.TopicSerial,
			Status:      item.Status,
			Title:       item.Title,
			AuthorName:  item.AuthorName,
			Description: item.Description,
			Tags:        tagMaps[item.Serial],
		})
	}

	return newsRes, nil
}

func (r *repository) CreateNews(news entity.News) error {
	return r.CreateNewsRepo(News{
		Serial:      news.Serial,
		TopicSerial: news.TopicSerial,
		AuthorName:  news.AuthorName,
		Status:      news.Status,
		Title:       news.Title,
		Description: news.Description,
	}, news.Tags)
}

func (r *repository) UpdateNews(news entity.News) error {
	return r.UpdateNewsRepo(News{
		Serial:      news.Serial,
		TopicSerial: news.TopicSerial,
		AuthorName:  news.AuthorName,
		Status:      news.Status,
		Title:       news.Title,
		Description: news.Description,
	}, news.Tags)
}
