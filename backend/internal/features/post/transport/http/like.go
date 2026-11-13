package post_transport

import (
	_ "actualiteco/actualite/internal/core/domain"
	"actualiteco/actualite/internal/core/logger"
	"actualiteco/actualite/internal/core/transport/response"
	httputils "actualiteco/actualite/internal/utils/http"
	"go.uber.org/zap"
	"net/http"
)

type LikeResponse struct {
	Message string `json:"message"`
}

// Like godoc
// @Summary Like a post
// @Description Like a post as the authenticated user, including their own posts.
// @Tags Posts
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {object} LikeResponse "Success"
// @Failure 400 {object} domain.CustomError "Invalid post ID or like state"
// @Failure 401 {object} domain.CustomError "Unauthorized"
// @Failure 404 {object} domain.CustomError "Post not found"
// @Failure 500 {object} domain.InternalError "Internal server error"
// @Router /posts/{postID}/like [post]
// @Security BearerAuth
func (h *PostsHTTPHandler) Like(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	respWriter := response.NewResponseHandler(log, w)
	userID, err := getUserUUIDFromContext(ctx)
	if err != nil {
		respWriter.MapError(err)
		return
	}
	postID, err := httputils.GetUUIDPathValue(r, "postID")
	if err != nil {
		respWriter.MapError(err)
		return
	}
	if err := h.PostsService.Like(ctx, userID, postID); err != nil {
		log.Error("like post", zap.Error(err))
		respWriter.MapError(err)
		return
	}
	respWriter.JSONResponse(LikeResponse{Message: "success"}, http.StatusOK)
}

// Unlike godoc
// @Summary Unlike a post
// @Description Unlike a post as the authenticated user, including their own posts.
// @Tags Posts
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {object} LikeResponse "Success"
// @Failure 400 {object} domain.CustomError "Invalid post ID or like state"
// @Failure 401 {object} domain.CustomError "Unauthorized"
// @Failure 404 {object} domain.CustomError "Post not found"
// @Failure 500 {object} domain.InternalError "Internal server error"
// @Router /posts/{postID}/unlike [post]
// @Security BearerAuth
func (h *PostsHTTPHandler) Unlike(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	respWriter := response.NewResponseHandler(log, w)
	userID, err := getUserUUIDFromContext(ctx)
	if err != nil {
		respWriter.MapError(err)
		return
	}
	postID, err := httputils.GetUUIDPathValue(r, "postID")
	if err != nil {
		respWriter.MapError(err)
		return
	}
	if err := h.PostsService.Unlike(ctx, userID, postID); err != nil {
		log.Error("unlike post", zap.Error(err))
		respWriter.MapError(err)
		return
	}
	respWriter.JSONResponse(LikeResponse{Message: "success"}, http.StatusOK)
}
