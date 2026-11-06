package auth_service

import (
	"actualiteco/actualite/internal/core/domain"
	hash_utils "actualiteco/actualite/internal/utils/hash"
	"actualiteco/actualite/internal/utils/jwt_utils"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (s *AuthService) LoginUser(ctx context.Context, email, password string) (*string, domain.User, error) {
	user, err := s.authRepo.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.User{}, fmt.Errorf("no user: %w", domain.ErrUnauthorized)
		}

		return nil, domain.User{}, fmt.Errorf("get from database: %w", domain.ErrInternal)
	}

	if !hash_utils.CheckPasswordHash(password, user.Password) {
		return nil, domain.User{}, fmt.Errorf("wrong password: %w", domain.ErrUnauthorized)
	}

	jwt, err := jwt_utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, domain.User{}, fmt.Errorf("generating jwt: %w", domain.ErrInternal)
	}

	return jwt, user, nil
}
