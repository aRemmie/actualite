package post_repository

import (
	postgres_pool "actualiteco/actualite/internal/core/repository/postgres"
	"context"
	"fmt"
	"uuid"
)

func (r *PostsRepository) GetCountOfReplies(ctx context.Context, parentID uuid.UUID) (uint64, bool, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, r.ConnPool.OpTimeout())
	defer cancel()

	conn := postgres_pool.ExecutorFromContext(ctx, r.ConnPool)

	query := `
		SELECT
			(SELECT COUNT(*) FROM posts WHERE parent_id = $1 AND deleted_at IS NULL) AS replies_count,
			EXISTS(SELECT 1 FROM posts WHERE id = $1 AND deleted_at IS NULL) AS post_exists;
	`

	row := conn.QueryRow(ctxTimeout, query, parentID)

	var count uint64
	var postExists bool
	err := row.Scan(&count, &postExists)
	if err != nil {
		return 0, false, fmt.Errorf("get count: %w", err)
	}

	return count, postExists, nil
}
