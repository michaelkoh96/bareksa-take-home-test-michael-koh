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

var (
	NewsTTL = (60 * 24)
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

	cacheKey := generateNewsCacheKey(status, topicSerialsArr)

	// Check cache
	cacheResult, err := h.cacheHelper.Get(cacheKey)
	if err != nil {
		log.Println(err.Error())
	}

	if cacheResult != nil {
		log.Println("cache hit")
		w.Write(cacheResult)
		return
	}

	// Get Topics
	var topics []entity.Topic
	if len(topicSerialsArr) > 0 {
		topics, err = h.topicService.GetTopicsBySerials(context.Background(), topicSerialsArr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Get News
	result, err := h.newsService.GetNews(context.Background(), entity.GetNewsQuery{
		Status:       status,
		TopicSerials: topicSerialsArr,
	})
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp = response.MapGetNewsResponse(result, topics)

	// Marshal response
	payload, err := json.Marshal(resp)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set cache
	err = h.cacheHelper.Set(cacheKey, payload, NewsTTL)
	if err != nil {
		log.Println(err.Error())
	}

	w.Write(payload)
}

func (h *handler) CreateNewsHandler(w http.ResponseWriter, r *http.Request) {
	var newNews entity.News

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newNews)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.newsService.CreateNews(context.Background(), newNews)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) UpdateNewsHandler(w http.ResponseWriter, r *http.Request) {
	var newNews entity.News

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newNews)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.newsService.UpdateNews(context.Background(), newNews)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteNewsHandler(w http.ResponseWriter, r *http.Request) {
	newsSerial := strings.TrimSpace(mux.Vars(r)["newsSerial"])
	if len(newsSerial) < 4 {
		http.Error(w, "Invalid news serial", http.StatusBadRequest)
		return
	}

	err := h.newsService.DeleteNews(context.Background(), newsSerial)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
