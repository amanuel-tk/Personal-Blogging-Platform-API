package repository

import (
	"context"
	"database/sql"

	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/model"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) CreatePost(ctx context.Context, post model.Post) (*model.Post, error) {

	query := `INSERT INTO post (title,content) VALUES ($1,$2) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, post.Title, post.Content).Scan(&post.ID)

	if err != nil {
		return nil, err
	}

	return &post, nil

}
