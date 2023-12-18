package login

import (
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
)

const (
	root = "/auth"
)

func NewHandlers(accountService services.IAccountService, authService services.IAuthService) []web.IHandler {
	return []web.IHandler{
		newPostHandler(accountService, authService),
	}
}
