package service

import (
	"context"
	"errors"
	"strings"

	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/model"
	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s PostService) CreatePost(ctx context.Context, post model.Post) (*model.Post, error) {

	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)

	if post.Title == "" {
		return nil, errors.New("title is required")
	}
	if post.Content == "" {
		return nil, errors.New("content is required")
	}
	return s.repo.CreatePost(ctx, post)
}

func (s PostService) GetPost(ctx context.Context) ([]model.Post, error) {
	return s.repo.GetPost(ctx)
}
