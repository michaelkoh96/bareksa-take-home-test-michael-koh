package handler

import (
	"bareksa-take-home-test-michael-koh/core/entity"
	"bareksa-take-home-test-michael-koh/handler/response"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (h *handler) GetNewsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp interface{}
	var topicSerialsArr []string

	queries := mux.Vars(r)
	status := strings.TrimSpace(queries["status"])
	topicSerials := strings.TrimSpace(queries["topicSerials"])

	if topicSerials == "" {
		topicSerialsArr = make([]string, 0)
	} else {
		topicSerialsArr = strings.Split(topicSerials, ",")
	}

	// Get News
	result, err := h.newsService.GetNews(context.Background(), entity.GetNewsQuery{
		Status:       status,
		TopicSerials: topicSerialsArr,
	})
	if err != nil {
		log.Println(err.Error())
		resp = response.MapErrorResponse(err)
	} else {
		resp = response.MapGetNewsResponse(result)
	}

	// Marshal response
	payload, err := json.Marshal(resp)
	if err != nil {
		log.Println(err.Error())
	}

	w.Write(payload)
}
