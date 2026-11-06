package profile_service

import (
	"actualiteco/actualite/internal/core/domain"
	"context"
	"fmt"
	"uuid"
)

func (s *ProfileService) Follow(ctx context.Context, followerID, followingID uuid.UUID) error {
	if followerID == followingID {
		return fmt.Errorf("user can't follow themselves: %w", domain.ErrInvalidRequest)
	}

	if err := s.repo.Follow(ctx, followerID, followingID); err != nil {
		return fmt.Errorf("follow user: %w", err)
	}

	return nil
}
