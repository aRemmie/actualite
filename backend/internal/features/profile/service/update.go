package profile_service

import (
	"actualiteco/actualite/internal/core/domain"
	"context"
	"fmt"
	"uuid"
)

// TODO: wrap in transaction
func (s *ProfileService) Update(ctx context.Context, id uuid.UUID, displayName, username, status *string) (domain.User, error) {
	if displayName == nil && username == nil && status == nil {
		return domain.User{}, fmt.Errorf("no fields to update: %w", domain.ErrInvalidRequest)
	}

	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	if displayName != nil {
		if err := user.UpdateDisplayName(*displayName); err != nil {
			return domain.User{}, fmt.Errorf("update display_name: %w", err)
		}
	}

	if username != nil {
		if err := user.UpdateUsername(*username); err != nil {
			return domain.User{}, fmt.Errorf("update username: %w", err)
		}
	}

	if status != nil {
		if err := user.UpdateStatus(*status); err != nil {
			return domain.User{}, fmt.Errorf("update status: %w", err)
		}
	}

	if err := s.repo.UpdateProfile(ctx, id, displayName, username, status); err != nil {
		return domain.User{}, fmt.Errorf("failed update user: %w", err)
	}

	updatedUser, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get updated user: %w", err)
	}

	return updatedUser, nil
}
