package post_repository

import (
	"actualiteco/actualite/internal/core/domain"
	postgres_pool "actualiteco/actualite/internal/core/repository/postgres"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"uuid"
)

func (r *PostsRepository) Unlike(ctx context.Context, userID, postID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.ConnPool.OpTimeout())
	defer cancel()
	executor := postgres_pool.ExecutorFromContext(ctx, r.ConnPool)
	var id string
	err := executor.QueryRow(ctx, `SELECT id FROM posts WHERE id = $1 AND deleted_at IS NULL FOR UPDATE`, postID.String()).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("post not found: %w", domain.ErrNotFound)
	}
	if err != nil {
		return fmt.Errorf("lock post: %w", err)
	}
	tag, err := executor.Exec(ctx, `DELETE FROM likes WHERE user_id = $1 AND post_id = $2`, userID.String(), postID.String())
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return fmt.Errorf("like reference not found: %w", domain.ErrNotFound)
		}
		return fmt.Errorf("unlike post: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("invalid like state: %w", domain.ErrInvalidRequest)
	}
	if _, err := executor.Exec(ctx, `UPDATE posts SET likes_count = GREATEST(likes_count - 1, 0) WHERE id = $1`, postID.String()); err != nil {
		return fmt.Errorf("update likes count: %w", err)
	}
	return nil
}
