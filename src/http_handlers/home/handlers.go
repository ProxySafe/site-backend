package home

import (
	"context"

	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
)

func NewHandlers(ctx context.Context, accountService services.IAccountService) []web.IHandler {
	return []web.IHandler{
		newHomeHandler(ctx, accountService),
	}
}
