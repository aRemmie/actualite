package profile_transport

import (
	"actualiteco/actualite/internal/core/domain"
	"actualiteco/actualite/internal/core/logger"
	"actualiteco/actualite/internal/core/transport/response"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// Update godoc
// @Summary Обновить профиль пользователя
// @Description Обновить профиль пользователя (display_name, username, status)
// @Tags Profile
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <jwt токен>"
// @Param request body UpdateProfileRequest true "Информация профиля"
// @Success 200 {object} ProfileResponse "Успешный запрос"
// @Failure 400 {object} domain.CustomError "Неверный запрос"
// @Failure 403 {object} domain.CustomError "Пользователь неавторизован"
// @Failure 409 {object} domain.CustomError "Пользователь с таким юзернеймом уже существует"
// @Failure 500 {object}  domain.InternalError "Внутренняя ошибка сервера"
// @Router /profile [patch]
func (h *ProfileHTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	respWriter := response.NewResponseHandler(log, w)

	userID, err := getUserUUIDFromContext(ctx)
	if err != nil {
		log.Error("get userID from context", zap.Error(err))
		respWriter.ErrorResponse(domain.ErrInternal, http.StatusInternalServerError)
		return
	}

	var req UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Debug("failed to decode request", zap.Error(err))
		respWriter.ErrorResponse(domain.ErrInvalidRequest, http.StatusBadRequest)
		return
	}

	user, err := h.service.Update(ctx, userID, req.DisplayName, req.Username, req.Status)
	if err != nil {
		log.Error("failed to update profile", zap.Error(err))
		respWriter.MapError(err)
		return
	}

	respWriter.JSONResponse(toResponse(user), http.StatusOK)
}
