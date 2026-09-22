package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/model"
	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/repository"
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

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})

		return
	}

	fmt.Println(post)

	createdPost, err := h.service.CreatePost(r.Context(), post)

	if err != nil {
		log.Printf("failed to create post :%v", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to create post"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)

}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	filters := repository.Filters{
		Title: r.URL.Query().Get("title"),
		Tags:  r.URL.Query()["tags"],
	}
	getPosts, err := h.service.GetPost(r.Context(), filters)

	if err != nil {
		log.Printf("failed to fetch posts: %v", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to fetch posts"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getPosts)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var req model.Post

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "id is required"})
		return
	}

	updatePost, err := h.service.UpdatePost(r.Context(), id, req)

	if err != nil {
		log.Printf("failed to update post:%v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to update post"})
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"message": "successfully updated", "post": updatePost})
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.DeletePost(r.Context(), id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": " post deleted successfully",
	})
}

func (h *PostHandler) GetSinglePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	result, err := h.service.GetSinglePost(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
