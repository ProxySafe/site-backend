package profile

import (
	"github.com/ProxySafe/site-backend/src/http_handlers/middleware"
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
)

const (
	root = "/profile"
)

func NewHandlers(proxyService services.IProxyService, authService services.IAuthService) []web.IHandler {
	return []web.IHandler{
		middleware.NewMiddleware(newProxiesHandler(proxyService), authService),
	}
}
