package handler

import (
	newsService "bareksa-take-home-test-michael-koh/core/service/news"
	tagService "bareksa-take-home-test-michael-koh/core/service/tag"
	topicService "bareksa-take-home-test-michael-koh/core/service/topic"
	"bareksa-take-home-test-michael-koh/pkg/cache"
	"net/http"
)

type (
	Handler interface {
		// News
		GetNewsHandler(w http.ResponseWriter, r *http.Request)
		CreateNewsHandler(w http.ResponseWriter, r *http.Request)
		UpdateNewsHandler(w http.ResponseWriter, r *http.Request)
		DeleteNewsHandler(w http.ResponseWriter, r *http.Request)

		// Tag
		GetTagsHandler(w http.ResponseWriter, r *http.Request)
		CreateTagsHandler(w http.ResponseWriter, r *http.Request)
		UpdateTagsHandler(w http.ResponseWriter, r *http.Request)
		DeleteTagsHandler(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		newsService  newsService.NewsService
		topicService topicService.TopicService
		tagService   tagService.TagsService
		cacheHelper  cache.CacheHelper
	}
)

func NewBareksaNewsHandler(newsService newsService.NewsService, topicService topicService.TopicService, tagService tagService.TagsService, cachehelper cache.CacheHelper) Handler {
	return &handler{
		newsService:  newsService,
		topicService: topicService,
		tagService:   tagService,
		cacheHelper:  cachehelper,
	}
}
