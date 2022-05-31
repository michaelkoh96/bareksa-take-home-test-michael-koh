package handler

import (
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
