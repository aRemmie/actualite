package post_repository

import (
	"actualiteco/actualite/internal/core/domain"
	postgres_pool "actualiteco/actualite/internal/core/repository/postgres"
	"time"
	"uuid"
)

type PostsRepository struct {
	ConnPool postgres_pool.Pool
}

func New(connPool postgres_pool.Pool) *PostsRepository {
	return &PostsRepository{
		ConnPool: connPool,
	}
}

type PostModel struct {
	ID         uuid.UUID  `db:"id"`
	UserID     uuid.UUID  `db:"user_id"`
	ParentID   *uuid.UUID `db:"parent_id"`
	Content    string     `db:"content"`
	ImageURL   *string    `db:"image_url"`
	LikesCount uint       `db:"likes_count"`
	DeletedAt  *time.Time `db:"deleted_at"`
	CreatedAt  time.Time  `db:"created_at"`
}

func toDomain(m PostModel) domain.Post {
	return domain.Post{
		ID:         m.ID,
		UserID:     m.UserID,
		ParentID:   m.ParentID,
		Content:    m.Content,
		ImageURL:   m.ImageURL,
		LikesCount: m.LikesCount,
		CreatedAt:  m.CreatedAt,
	}
}
