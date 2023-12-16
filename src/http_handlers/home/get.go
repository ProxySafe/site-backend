package home

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
)

type homeHandler struct {
	method string

	accountService services.IAccountService
}

func newHomeHandler(accountService services.IAccountService) web.IHandler {
	return &homeHandler{
		method:         http.MethodGet,
		accountService: accountService,
	}
}

func (h *homeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accounts, err := h.accountService.FindAll(r.Context())
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(accounts[0])
}

func (h *homeHandler) GetMethod() string {
	return h.method
}

func (h *homeHandler) GetPath() string {
	return "/"
}
