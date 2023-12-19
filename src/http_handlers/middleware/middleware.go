package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
)

const (
	authorizationHeader = "Authorization"
)

type middleware struct {
	handler     web.IHandler
	authService services.IAuthService
}

type responseDto struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewMiddleware(handler web.IHandler, authService services.IAuthService) web.IHandler {
	return &middleware{
		handler:     handler,
		authService: authService,
	}
}

func (m *middleware) GetMethod() string {
	return m.handler.GetMethod()
}

func (m *middleware) GetPath() string {
	return m.handler.GetPath()
}

func (m *middleware) Handle(w http.ResponseWriter, r *http.Request) {
	resp := &responseDto{}
	_, valid, err := m.authService.ParseToken(r.Context(), r.Header.Get(authorizationHeader))
	if err != nil || !valid {
		resp.Message = err.Error()
		resp.Status = http.StatusUnauthorized
		json.NewEncoder(w).Encode(resp)
		return
	}

	m.handler.Handle(w, r)
}
