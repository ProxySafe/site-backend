package home

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
)

type homeHandler struct {
	ctx    context.Context
	method string

	accountService services.IAccountService
}

func newHomeHandler(ctx context.Context, accountService services.IAccountService) web.IHandler {
	return &homeHandler{
		ctx:            ctx,
		method:         http.MethodGet,
		accountService: accountService,
	}
}

func (h *homeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accounts, err := h.accountService.FindAll(h.ctx)
	if err != nil {
		log.Fatal(err)
	}
	// resp := &responses.HomeResponseDTO{
	// 	UserId: uuid.NewString(),
	// }
	json.NewEncoder(w).Encode(accounts[0])
}

func (h *homeHandler) GetMethod() string {
	return h.method
}

func (h *homeHandler) GetPath() string {
	return "/"
}
