package post_service

import (
	"actualiteco/actualite/internal/core/domain"
	"context"
	"fmt"
	"uuid"
)

func (s *PostsService) GetPosts(ctx context.Context, userID uuid.UUID, page, pageSize *uint64) ([]domain.PostWithAuthor, domain.Pagination, error) {
	pagination := domain.NewPagination(page, pageSize)
	offset, limit, err := pagination.GetOffsetLimit()
	if err != nil {
		return nil, domain.Pagination{}, fmt.Errorf("get offset and limit: %w", err)
	}

	posts, count, err := s.postsRepo.GetPostsFeed(ctx, userID, limit, offset)
	if err != nil {
		return nil, domain.Pagination{}, fmt.Errorf("get posts feed: %w", err)
	}

	pagination.ParseCount(count)

	return posts, pagination, nil
}
