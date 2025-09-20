package store

import (
	"github.com/ProxySafe/site-backend/src/http_handlers/middleware"
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
)

const (
	root = "/store"
)

func NewHandlers(
	proxyService services.IProxyService,
	orderService services.IOrderService,
	authService services.IAuthService,
) []web.IHandler {
	return []web.IHandler{
		middleware.NewMiddleware(newBuyHandler(proxyService, orderService), authService),
	}
}
