package auth

import (
	"net/http"

	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
)

type refreshHandler struct {
	accountService services.IAccountService
	authService    services.IAuthService
}

type refreshRequestDto struct {
}

type refreshResponseDto struct{}

func newRefreshHandler(accountService services.IAccountService, authService services.IAuthService) web.IHandler {
	return &refreshHandler{
		accountService: accountService,
		authService:    authService,
	}
}

func (h *refreshHandler) Handle(w http.ResponseWriter, r *http.Request) {

}

func (h *refreshHandler) GetMethod() string {
	return http.MethodGet
}

func (h *refreshHandler) GetPath() string {
	return root + "/refresh/"
}
