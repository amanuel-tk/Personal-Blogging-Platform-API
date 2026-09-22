package model

type Post struct {
	ID        string   `json:"id"`
	Title     string   `json:"title" validate:"required,min=3,max=100"`
	Category  string   `json:"category" validate:"required,min=3,max=100"`
	Content   string   `json:"content" validate:"required,min=3,max=200"`
	Tags      []string `json:"tags" validate:"required"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}
