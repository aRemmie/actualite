package post_service

import (
	"actualiteco/actualite/internal/core/domain"
	"context"
	"fmt"
	"uuid"
)

func (s *PostsService) GetReplies(ctx context.Context, postID uuid.UUID, page, pageSize *uint64) ([]domain.PostWithAuthor, domain.Pagination, error) {
	pagination := domain.NewPagination(page, pageSize)
	offset, limit, err := pagination.GetOffsetLimit()
	if err != nil {
		return nil, domain.Pagination{}, fmt.Errorf("get offset and limit: %w", err)
	}

	posts := make([]domain.PostWithAuthor, 0)
	err = s.txManager.Do(ctx, func(ctx context.Context) error {
		count, postExists, err := s.postsRepo.GetCountOfReplies(ctx, postID)
		if err != nil {
			return fmt.Errorf("get count of posts: %w", err)
		}
		if !postExists {
			return fmt.Errorf("no post with id=%s: %w", postID.String(), domain.ErrNotFound)
		}

		pagination.ParseCount(count)
		if count == 0 {
			return nil
		}

		posts, err = s.postsRepo.GetReplies(ctx, postID, limit, offset)
		if err != nil {
			return fmt.Errorf("get replies: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, domain.Pagination{}, err
	}

	return posts, pagination, nil
}
