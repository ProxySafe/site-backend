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

type refreshHandler struct {
	authService services.IAuthService
}

type refreshRequestDto struct {
	Fingerprint    int    `json:"fingerprint"`
	Os             string `json:"os"`
	UserAgent      string `json:"user_agent"`
	OldAccessToken string `json:"old_access_token"`
	RefreshToken   string `json:"refresh_token"`
}

type refreshResponseDto struct {
	common.StandardResponse
	AccessToken string `json:"accessToken"`
}

func newRefreshHandler(authService services.IAuthService) web.IHandler {
	return &refreshHandler{
		authService: authService,
	}
}

func (h *refreshHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &refreshRequestDto{}
	resp := &refreshResponseDto{}

	defer func() {
		common.EnableCors(w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}()

	if err := utils.SetRequestDto(r, req); err != nil {
		resp.SetError(err, http.StatusBadRequest)
		return
	}

	newAccessToken, err := h.authService.RefreshAccessToken(
		r.Context(),
		req.OldAccessToken,
		req.RefreshToken,
		entities.Fingerprint{
			Fingerprint: int64(req.Fingerprint),
			Os:          req.Os,
			UserAgent:   req.UserAgent,
		},
	)
	if err != nil {
		resp.SetError(err, http.StatusUnauthorized)
		return
	}

	resp.SetError(nil, http.StatusOK)
	resp.AccessToken = newAccessToken
}

func (h *refreshHandler) GetMethod() string {
	return http.MethodPost
}

func (h *refreshHandler) GetPath() string {
	return root + "/refresh/"
}
