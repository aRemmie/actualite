package httputils

import (
	"actualiteco/actualite/internal/core/domain"
	"fmt"
	"net/http"
	"uuid"
)

func GetUUIDPathValue(r *http.Request, key string) (uuid.UUID, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return uuid.Nil(), fmt.Errorf(
			"no key='%s' in path values: %w",
			key,
			domain.ErrInvalidRequest,
		)
	}

	id, err := uuid.Parse(pathValue)
	if err != nil {
		return uuid.Nil(), fmt.Errorf(
			"failed to parse path value='%s' by key='%s' to uuid: %w: %w",
			pathValue,
			key,
			err,
			domain.ErrInvalidRequest,
		)
	}

	return id, nil
}
