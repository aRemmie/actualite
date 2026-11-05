package post_service

import (
	"actualiteco/actualite/internal/core/domain"
	"context"
)

type PostsService struct {
	postsRepo PostsRepository
}

type PostsRepository interface {
	CreatePost(ctx context.Context, post domain.Post) (domain.Post, error)
}

func New(postsRepo PostsRepository) *PostsService {
	return &PostsService{
		postsRepo: postsRepo,
	}
}
