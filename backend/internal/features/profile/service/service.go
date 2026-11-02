package profile_service

import (
	"actualiteco/actualite/internal/core/domain"
	"context"
	"uuid"
)

type ProfileService struct {
	repo usersRepo
}

type usersRepo interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (domain.User, error)
}

func NewProfileService(repo usersRepo) *ProfileService {
	return &ProfileService{
		repo: repo,
	}
}
