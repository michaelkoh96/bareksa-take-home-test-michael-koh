package handler

import (
	"bareksa-take-home-test-michael-koh/core/entity"
	"bareksa-take-home-test-michael-koh/handler/response"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (h *handler) GetTagsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queries := mux.Vars(r)
	pageInt, err := strconv.Atoi(strings.TrimSpace(queries["page"]))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sizeInt, err := strconv.Atoi(strings.TrimSpace(queries["size"]))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cacheKey := generateTagsCacheKey(pageInt, sizeInt)

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

	resTags, err := h.tagService.GetTags(context.Background(), pageInt, sizeInt)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := response.MapGetTagResponse(resTags)

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

	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func (h *handler) DeleteTagsHandler(w http.ResponseWriter, r *http.Request) {
	tagName := strings.TrimSpace(mux.Vars(r)["tagName"])
	if len(tagName) == 0 {
		http.Error(w, "Invalid tag name", http.StatusBadRequest)
		return
	}

	err := h.tagService.DeleteTag(context.Background(), tagName)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateTagsHandler(w http.ResponseWriter, r *http.Request) {
	var tag entity.Tag
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tag)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.tagService.CreateTags(context.Background(), tag.Name)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) UpdateTagsHandler(w http.ResponseWriter, r *http.Request) {
	tagName := strings.TrimSpace(mux.Vars(r)["tagName"])
	if len(tagName) == 0 {
		http.Error(w, "Invalid tag name", http.StatusBadRequest)
		return
	}

	var tag entity.Tag
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tag)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.tagService.UpdateTags(context.Background(), tagName, tag.Name)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
