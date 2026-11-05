package post_transport

import (
	"actualiteco/actualite/internal/core/domain"
	"time"
	"uuid"
)

type CreatePostDTO struct {
	Content  string     `json:"content" example:"Hello, world!" validate:"required,min=1,max=280"`
	ParentID *uuid.UUID `json:"parent_id" example:"123e4567-e89b-12d3-a456-426614174000" validate:"omitempty"`
	ImageURL *string    `json:"image_url" example:"https://example.com/image.jpg" validate:"omitempty,url"`
}

type CreatePostDTOResponse struct {
	ID         uuid.UUID  `json:"id" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	UserID     uuid.UUID  `json:"user_id" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	ParentID   *uuid.UUID `json:"parent_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Content    string     `json:"content" example:"Hello, world!"`
	ImageURL   *string    `json:"image_url" example:"https://example.com/image.jpg"`
	LikesCount uint       `json:"likes_count" example:"10"`
	CreatedAt  time.Time  `json:"created_at" example:"2026-03-25T12:00:41.267Z"`
}

func toResponse(p domain.Post) CreatePostDTOResponse {
	return CreatePostDTOResponse{
		ID:         p.ID,
		UserID:     p.UserID,
		ParentID:   p.ParentID,
		Content:    p.Content,
		ImageURL:   p.ImageURL,
		LikesCount: p.LikesCount,
		CreatedAt:  p.CreatedAt,
	}
}
