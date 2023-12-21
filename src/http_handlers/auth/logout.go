package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/http_handlers/common"
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
	"github.com/ProxySafe/site-backend/src/utils"
)

type logoutHandler struct {
	authService services.IAuthService
}

type logoutRequestDto struct {
	// TODO: fingerprint?
	AccessToken string `json:"access_token"`
	entities.Fingerprint
}

type logoutResponseDto struct {
	common.StandardResponse
}

func newLogoutHandler(authService services.IAuthService) web.IHandler {
	return &logoutHandler{
		authService: authService,
	}
}

func (h *logoutHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &logoutRequestDto{}
	resp := &logoutResponseDto{}

	defer func() {
		common.EnableCors(w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}()

	if err := utils.SetRequestDto(r, req); err != nil {
		resp.SetError(err, http.StatusBadRequest)
		return
	}

	if err := h.authService.RemoveRefreshToken(r.Context(), req.AccessToken, &req.Fingerprint); err != nil {
		resp.SetError(err, http.StatusUnauthorized)
	}
	resp.SetError(nil, http.StatusOK)
}

func (h *logoutHandler) GetMethod() string {
	return http.MethodPost
}

func (h *logoutHandler) GetPath() string {
	return root + "/logout/"
}
