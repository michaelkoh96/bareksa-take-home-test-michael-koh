package handler

import (
	newsService "bareksa-take-home-test-michael-koh/core/service/news"
	"net/http"
)

type (
	Handler interface {
		GetNewsHandler(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		newsService newsService.NewsService
	}
)

func NewBareksaNewsHandler(newsService newsService.NewsService) Handler {
	return &handler{
		newsService: newsService,
	}
}