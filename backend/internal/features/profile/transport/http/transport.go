package profile_transport

import (
	"actualiteco/actualite/internal/core/domain"
	"actualiteco/actualite/internal/core/middleware"
	"actualiteco/actualite/internal/core/transport/server"
	"context"
	"net/http"
	"uuid"
)

type ProfileHTTPHandler struct {
	service profileService
}

type profileService interface {
	GetProfile(ctx context.Context, id uuid.UUID) (domain.User, error)
}

func NewProfileHTTPHandler(serv profileService) *ProfileHTTPHandler {
	return &ProfileHTTPHandler{
		service: serv,
	}
}

func (h *ProfileHTTPHandler) Routes() []server.Route {
	return []server.Route{
		{
			Method:               http.MethodGet,
			URL:                  "/profile/me",
			Handler:              h.GetUserProfileByJWT,
			AdditionalMiddleware: []middleware.Middleware{middleware.AuthMiddleware()},
		},
	}
}
