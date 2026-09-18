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

	query := `INSERT INTO posts (title,content) VALUES ($1,$2) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, post.Title, post.Content).Scan(&post.ID)

	if err != nil {
		return nil, err
	}

	return &post, nil

}

func (r *PostRepository) GetPost(ctx context.Context) ([]model.Post, error) {
	query := `SELECT id,title,content,COALESCE(tags,''),created_at,updated_at FROM posts`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post

	for rows.Next() {
		var post model.Post

		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Tags,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil

}
