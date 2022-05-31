package handler

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

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
