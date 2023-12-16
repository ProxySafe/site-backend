package home

import (
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
)

func NewHandlers(accountService services.IAccountService) []web.IHandler {
	return []web.IHandler{
		newHomeHandler(accountService),
	}
}
