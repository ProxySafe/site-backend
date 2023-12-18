package login

import (
	"net/http"

	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
)

type postHandler struct {
	accountService services.IAccountService
	authService    services.IAuthService
}

func newPostHandler(accountService services.IAccountService, authService services.IAuthService) web.IHandler {
	return &postHandler{
		accountService: accountService,
	}
}

func (h *postHandler) Handle(w http.ResponseWriter, r *http.Request) {
	userName := r.Header.Get("User-Name")
	hashedPassword := r.Header.Get("Hashed-Password")
	if len(userName) == 0 || len(hashedPassword) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	account, err := h.accountService.GetByUsername(r.Context(), userName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if account.HashedPassword != hashedPassword {
		w.WriteHeader(http.StatusForbidden)
	}

}

func (h *postHandler) GetMethod() string {
	return http.MethodPost
}

func (h *postHandler) GetPath() string {
	return root + "/login/"
}
