package post_repository

import (
	"context"
	"fmt"
	"uuid"
)

func (r *PostsRepository) GetCountOfReplies(ctx context.Context, parentID uuid.UUID) (uint64, bool, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, r.ConnPool.OpTimeout())
	defer cancel()

	query := `
		SELECT
			(SELECT COUNT(*) FROM posts WHERE parent_id = $1) AS replies_count,
			EXISTS(SELECT 1 FROM posts WHERE id = $1) AS post_exists;
	`

	row := r.ConnPool.QueryRow(ctxTimeout, query, parentID)

	var count uint64
	var postExists bool
	err := row.Scan(&count, &postExists)
	if err != nil {
		return 0, false, fmt.Errorf("get count: %w", err)
	}

	return count, postExists, nil
}
