package handler

import (
	newsService "bareksa-take-home-test-michael-koh/core/service/news"
	tagService "bareksa-take-home-test-michael-koh/core/service/tag"
	topicService "bareksa-take-home-test-michael-koh/core/service/topic"
	"net/http"
)

type (
	Handler interface {
		GetNewsHandler(w http.ResponseWriter, r *http.Request)
		CreateNewsHandler(w http.ResponseWriter, r *http.Request)
		UpdateNewsHandler(w http.ResponseWriter, r *http.Request)
		DeleteNewsHandler(w http.ResponseWriter, r *http.Request)
		DeleteTagsHandler(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		newsService  newsService.NewsService
		topicService topicService.TopicService
		tagService   tagService.NewsService
	}
)

func NewBareksaNewsHandler(newsService newsService.NewsService, topicService topicService.TopicService, tagService tagService.NewsService) Handler {
	return &handler{
		newsService:  newsService,
		topicService: topicService,
		tagService:   tagService,
	}
}
