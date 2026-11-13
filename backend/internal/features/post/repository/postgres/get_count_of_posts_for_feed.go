package post_repository

import (
	postgres_pool "actualiteco/actualite/internal/core/repository/postgres"
	"context"
	"fmt"
	"uuid"
)

func (r *PostsRepository) GetCountOfPostsForFeed(ctx context.Context, userID uuid.UUID) (uint64, bool, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, r.ConnPool.OpTimeout())
	defer cancel()

	conn := postgres_pool.ExecutorFromContext(ctx, r.ConnPool)

	query := `
		WITH user_follows AS (
			SELECT following_id AS author_id
			FROM follows
			WHERE follower_id = $1
		)
		SELECT
			(SELECT COUNT(*)
			FROM posts p
			JOIN users u              ON u.id         = p.user_id
			LEFT JOIN user_follows uf ON uf.author_id = p.user_id
			WHERE p.parent_id IS NULL
				AND p.deleted_at IS NULL
				AND (uf.author_id IS NOT NULL
					OR p.created_at >= NOW() - INTERVAL '1 month')
			) AS total_count,
			EXISTS(SELECT 1 FROM users WHERE id = $1) AS user_exists;
	`

	row := conn.QueryRow(ctxTimeout, query, userID)

	var count uint64
	var userExists bool
	if err := row.Scan(&count, &userExists); err != nil {
		return 0, false, fmt.Errorf("scan row: %w", err)
	}

	return count, userExists, nil
}
