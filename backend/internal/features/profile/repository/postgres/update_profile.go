package users_repository

import (
	"actualiteco/actualite/internal/core/domain"
	"context"
	"errors"
	"fmt"
	"uuid"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
)

const uniqueViolation = "23505"

func (r *UsersRepository) UpdateProfile(ctx context.Context, id uuid.UUID, displayName, username, status *string) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, r.ConnPool.OpTimeout())
	defer cancel()

	query := sq.
		Update("users")

	if displayName != nil {
		query = query.Set("display_name", displayName)
	}

	if username != nil {
		query = query.Set("username", username)
	}

	if status != nil {
		query = query.Set("status", status)
	}

	sqlQuery, args, err := query.Where(sq.Eq{"id": id.String()}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("building sql query: %w", err)
	}

	tag, err := r.ConnPool.Exec(ctxTimeout, sqlQuery, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == uniqueViolation {
			return fmt.Errorf("user with this username already exists: %w", domain.ErrConflict)
		}
		return fmt.Errorf("exec query: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("no updated rows: %w", domain.ErrNotFound)
	}

	return nil
}
