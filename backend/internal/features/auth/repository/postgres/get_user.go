package users_repository

import (
	"actualiteco/actualite/internal/core/domain"
	auth_models "actualiteco/actualite/internal/features/auth/repository"
	"context"
	"fmt"
)

func (r *UsersRepository) GetUser(ctx context.Context, email string) (domain.User, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, r.ConnPool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, email, password, role, created_at, username, display_name, status, avatar_url, email_verified
	FROM users
	WHERE email=$1
	`

	row := r.ConnPool.QueryRow(ctxTimeout, query, email)

	var model auth_models.UserModel
	err := row.Scan(&model.ID, &model.Email, &model.Password, &model.Role, &model.CreatedAt,
		&model.Username, &model.DisplayName, &model.Status, &model.AvatarURL, &model.EmailVerified)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan returning user: %w", err)
	}

	return auth_models.ToDomain(model), nil
}
