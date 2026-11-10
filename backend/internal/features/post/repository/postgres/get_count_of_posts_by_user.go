package post_repository

import (
	"context"
	"fmt"
	"uuid"
)

func (r *PostsRepository) GetCountOfPostsByUserID(ctx context.Context, userID uuid.UUID) (uint64, bool, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, r.ConnPool.OpTimeout())
	defer cancel()

	query := `
		SELECT
			(SELECT COUNT(*) FROM posts WHERE user_id = $1) AS posts_count,
			EXISTS(SELECT 1 FROM users WHERE id = $1) AS user_exists;
	`

	row := r.ConnPool.QueryRow(ctxTimeout, query, userID)

	var count uint64
	var userExists bool
	err := row.Scan(&count, &userExists)
	if err != nil {
		return 0, false, fmt.Errorf("get count: %w", err)
	}

	return count, userExists, nil
}
