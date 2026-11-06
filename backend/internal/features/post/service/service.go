package post_service

import (
	"actualiteco/actualite/internal/core/domain"
	"context"
	"uuid"
)

type PostsService struct {
	postsRepo PostsRepository
}

type PostsRepository interface {
	CreatePost(ctx context.Context, post domain.Post) (domain.Post, error)
	GetPostByID(ctx context.Context, postID uuid.UUID) (domain.PostWithAuthor, error)
}

func New(postsRepo PostsRepository) *PostsService {
	return &PostsService{
		postsRepo: postsRepo,
	}
}
