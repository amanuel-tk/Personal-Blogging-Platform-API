package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/model"
	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/repository"
	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/service"
	"github.com/go-playground/validator/v10"
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

		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	createdPost, err := h.service.CreatePost(r.Context(), post)

	if err != nil {

		var validationError validator.ValidationErrors

		if errors.As(err, &validationError) {
			writeError(w, http.StatusBadRequest, "invalid post data")
			return
		}
		log.Printf("failed to create post :%v", err)

		writeError(w, http.StatusInternalServerError, "failed to create post")
		return
	}

	writeJson(w, http.StatusCreated, createdPost)

}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	filters := repository.Filters{
		Title: r.URL.Query().Get("title"),
		Tags:  r.URL.Query()["tags"],
	}
	getPosts, err := h.service.GetPost(r.Context(), filters)

	if err != nil {
		log.Printf("failed to fetch posts: %v", err)

		writeError(w, http.StatusInternalServerError, "failed to fetch posts")
		return
	}

	writeJson(w, http.StatusOK, getPosts)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var req model.Post

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		writeError(w, http.StatusBadRequest, "bad request body")
		return
	}

	updatePost, err := h.service.UpdatePost(r.Context(), id, req)

	if err != nil {
		var validationErrors validator.ValidationErrors
		log.Printf("error %v", err)
		switch {
		case errors.Is(err, service.ErrIdIsRequired):
			writeError(w, http.StatusBadRequest, service.ErrIdIsRequired.Error())

		case errors.Is(err, service.ErrContentMissing):
			writeError(w, http.StatusBadRequest, service.ErrContentMissing.Error())

		case errors.As(err, &validationErrors):
			writeError(w, http.StatusBadRequest, "invalid post data")

		case errors.Is(err, sql.ErrNoRows):
			writeError(w, http.StatusNotFound, "post not found")

		default:
			writeError(w, http.StatusInternalServerError, "failed to update post")
		}

		return
	}
	data := map[string]any{"message": "successfully updated", "post": updatePost}

	writeJson(w, http.StatusOK, data)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.DeletePost(r.Context(), id)

	if err != nil {

		switch {
		case errors.Is(err, service.ErrIdIsRequired):
			writeError(w, http.StatusBadRequest, service.ErrIdIsRequired.Error())
		case errors.Is(err, repository.ErrPostNotFound):
			writeError(w, http.StatusNotFound, repository.ErrPostNotFound.Error())
		default:
			writeError(w, http.StatusInternalServerError, "failed to delete post")
		}
		return
	}

	writeJson(w, http.StatusNoContent, nil)
}

func (h *PostHandler) GetSinglePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	result, err := h.service.GetSinglePost(r.Context(), id)
	if err != nil {

		switch {
		case errors.Is(err, service.ErrIdIsRequired):
			writeError(w, http.StatusBadRequest, service.ErrIdIsRequired.Error())
		case errors.Is(err, sql.ErrNoRows):
			writeError(w, http.StatusNotFound, "post not found")
		default:
			writeError(w, http.StatusInternalServerError, "failed to delete post")
		}
		return
	}

	writeJson(w, http.StatusOK, result)

}
