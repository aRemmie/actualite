package post_service

import (
	"context"
	"fmt"
	"uuid"
)

func (s *PostsService) Like(ctx context.Context, userID, postID uuid.UUID) error {
	return s.txManager.DoWrite(ctx, func(ctx context.Context) error {
		if err := s.postsRepo.Like(ctx, userID, postID); err != nil {
			return fmt.Errorf("like post: %w", err)
		}
		return nil
	})
}
func (s *PostsService) Unlike(ctx context.Context, userID, postID uuid.UUID) error {
	return s.txManager.DoWrite(ctx, func(ctx context.Context) error {
		if err := s.postsRepo.Unlike(ctx, userID, postID); err != nil {
			return fmt.Errorf("unlike post: %w", err)
		}
		return nil
	})
}
