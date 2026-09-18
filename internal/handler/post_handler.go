package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/model"
	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/service"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println(post)

	createdPost, err := h.service.CreatePost(r.Context(), post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)

}

func (h PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	getPosts, err := h.service.GetPost(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getPosts)
}
