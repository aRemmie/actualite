package post_transport

import (
	"actualiteco/actualite/internal/core/domain"
	"actualiteco/actualite/internal/core/middleware"
	"actualiteco/actualite/internal/core/transport/server"
	"context"
	"net/http"
	"uuid"
)

type PostsHTTPHandler struct {
	PostsService PostsService
}

type PostsService interface {
	Create(ctx context.Context, userID uuid.UUID, content string, parentID *uuid.UUID, imageURL *string) (domain.Post, error)
}

func NewPostsHandler(postsService PostsService) *PostsHTTPHandler {
	return &PostsHTTPHandler{
		PostsService: postsService,
	}
}

func (h *PostsHTTPHandler) Routes() []server.Route {
	return []server.Route{
		{
			Method:               http.MethodPost,
			URL:                  "/posts/create",
			Handler:              h.Create,
			AdditionalMiddleware: []middleware.Middleware{middleware.AuthMiddleware()},
		},
	}
}
