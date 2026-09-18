package repository

import (
	"context"
	"database/sql"
	"errors"

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

func (r *PostRepository) UpdatePost(ctx context.Context, id string, updatedPost model.Post) (*model.Post, error) {

	query := `UPDATE posts SET title=$1,content=$2,tags=$3,updated_at=CURRENT_TIMESTAMP WHERE id=$4 RETURNING title,content,COALESCE(tags,''),created_at,updated_at`

	err := r.db.QueryRowContext(ctx, query, updatedPost.Title, updatedPost.Content, updatedPost.Tags, id).Scan(&updatedPost.Title, &updatedPost.Content, &updatedPost.Tags, &updatedPost.CreatedAt, &updatedPost.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &updatedPost, nil

}

func (r *PostRepository) DeletePost(ctx context.Context, id string) error {
	query := `DELETE FROM posts WHERE id=$1`

	result, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("post not found")
	}
	return nil

}
