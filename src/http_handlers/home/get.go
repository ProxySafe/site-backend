package home

import (
	"fmt"
	"net/http"

	"github.com/ProxySafe/site-backend/src/modules/web"
)

type homeHandler struct {
	method string
}

func newHomeHandler() web.IHandler {
	return &homeHandler{
		method: http.MethodGet,
	}
}

func (h *homeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to ProxySafe!")
}

func (h *homeHandler) GetMethod() string {
	return h.method
}

func (h *homeHandler) GetPath() string {
	return "/"
}
