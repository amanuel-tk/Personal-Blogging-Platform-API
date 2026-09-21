package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/model"
	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/repository"
	"github.com/go-playground/validator/v10"
)

type PostService struct {
	repo      *repository.PostRepository
	validator *validator.Validate
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *PostService) CreatePost(ctx context.Context, post model.Post) (*model.Post, error) {

	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
	post.Tags = strings.TrimSpace(post.Tags)

	fmt.Println(post.Content)

	if err := s.validator.Struct(post); err != nil {
		return nil, err
	}

	return s.repo.CreatePost(ctx, post)
}

func (s *PostService) GetPost(ctx context.Context, filters repository.Filters) ([]model.Post, error) {
	return s.repo.GetPost(ctx, filters)
}

func (s *PostService) UpdatePost(ctx context.Context, id string, updatedPost model.Post) (*model.Post, error) {

	id = strings.TrimSpace(id)

	updatedPost.Content = strings.TrimSpace(updatedPost.Content)
	updatedPost.Title = strings.TrimSpace(updatedPost.Title)
	updatedPost.Tags = strings.TrimSpace(updatedPost.Tags)

	if err := s.validator.Struct(updatedPost); err != nil {
		return nil, err
	}

	if id == "" {
		return nil, errors.New("id is required")
	}

	if updatedPost.Title == "" || updatedPost.Content == "" || updatedPost.Tags == "" {
		return nil, errors.New("title, content and tags are required.")
	}

	return s.repo.UpdatePost(ctx, id, updatedPost)

}

func (s *PostService) DeletePost(ctx context.Context, id string) error {
	id = strings.TrimSpace(id)

	if id == "" {
		return errors.New("id is required")
	}

	return s.repo.DeletePost(ctx, id)
}

func (s *PostService) GetSinglePost(ctx context.Context, id string) (*model.Post, error) {
	id = strings.TrimSpace(id)
	if id == "" {
		return nil, errors.New("post id is required")
	}
	return s.repo.GetSinglePost(ctx, id)
}
