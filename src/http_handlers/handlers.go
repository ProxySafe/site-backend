package http_handlers

import (
	"github.com/ProxySafe/site-backend/src/http_handlers/middleware"
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
)

func WithMiddleware(authService services.IAuthService, handlers ...web.IHandler) []web.IHandler {
	withMiddlewareHandlers := make([]web.IHandler, 0, len(handlers))
	for _, handler := range handlers {
		withMiddlewareHandlers = append(withMiddlewareHandlers, middleware.NewMiddleware(handler, authService))
	}
	return withMiddlewareHandlers
}
