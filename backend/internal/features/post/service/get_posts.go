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

	posts := make([]domain.PostWithAuthor, 0)
	err = s.txManager.Do(ctx, func(ctx context.Context) error {
		count, userExists, err := s.postsRepo.GetCountOfPostsForFeed(ctx, userID)
		if err != nil {
			return fmt.Errorf("get count of posts for feed: %w", err)
		}

		if !userExists {
			return fmt.Errorf("no user with id=%s: %w", userID.String(), domain.ErrNotFound)
		}

		pagination.ParseCount(count)

		if count == 0 {
			return nil
		}

		posts, err = s.postsRepo.GetPostsFeed(ctx, userID, limit, offset)
		if err != nil {
			return fmt.Errorf("get posts feed: %w", err)
		}

		return nil
	})
	if err != nil {
		return []domain.PostWithAuthor{}, pagination, err
	}

	return posts, pagination, nil
}
