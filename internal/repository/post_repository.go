package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/amanuel-tk/Personal-Blogging-Platform-API/internal/model"
	"github.com/lib/pq"
)

type PostRepository struct {
	db *sql.DB
}

type Filters struct {
	Title string
	Tags  []string
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

var ErrPostNotFound = errors.New("post not found")

func (r *PostRepository) CreatePost(ctx context.Context, post model.Post) (*model.Post, error) {

	query := `INSERT INTO posts (title,category,content,tags) VALUES ($1,$2,$3,$4) RETURNING id,title,category,content,tags,created_at,updated_at`

	err := r.db.QueryRowContext(ctx, query, post.Title, post.Category, post.Content, pq.Array(post.Tags)).Scan(&post.ID, &post.Title, &post.Category, &post.Content, pq.Array(&post.Tags), &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &post, nil

}

func (r *PostRepository) GetPost(ctx context.Context, filters Filters) ([]model.Post, error) {
	baseQuery := `SELECT id,title,category,content,tags,created_at,updated_at FROM posts WHERE 1=1 `

	var args []any

	if filters.Title != "" {

		args = append(args, filters.Title)
		baseQuery += fmt.Sprintf("AND title = $%d ", len(args))

	}

	if len(filters.Tags) != 0 {
		args = append(args, pq.Array(filters.Tags))
		baseQuery += fmt.Sprintf("AND tags @> $%d::text[] ", len(args))
	}

	rows, err := r.db.QueryContext(ctx, baseQuery, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]model.Post, 0)

	for rows.Next() {
		var post model.Post

		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Category,
			&post.Content,
			pq.Array(&post.Tags),
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

	query := `UPDATE posts SET title=$1,category=$2,content=$3,tags=$4,updated_at=CURRENT_TIMESTAMP WHERE id=$5 RETURNING id,title,category,content,tags,created_at,updated_at`

	err := r.db.QueryRowContext(ctx, query, updatedPost.Title, updatedPost.Category, updatedPost.Content, pq.Array(updatedPost.Tags), id).
		Scan(&updatedPost.ID, &updatedPost.Title, &updatedPost.Category, &updatedPost.Content, pq.Array(&updatedPost.Tags), &updatedPost.CreatedAt, &updatedPost.UpdatedAt)

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
		return ErrPostNotFound
	}
	return nil

}

func (r *PostRepository) GetSinglePost(ctx context.Context, id string) (*model.Post, error) {

	query := `SELECT id,title,category,content,tags,created_at,updated_at FROM posts WHERE id=$1`

	var data model.Post

	err := r.db.QueryRowContext(ctx, query, id).Scan(&data.ID, &data.Title, &data.Category, &data.Content, pq.Array(&data.Tags), &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
