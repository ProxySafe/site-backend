package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ProxySafe/site-backend/src/http_handlers/common"
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
	"github.com/ProxySafe/site-backend/src/utils"
)

type registerHandler struct {
	accountService services.IAccountService
	emailService   services.IEmailService
}

type registerHandlerRequestDto struct {
	AccountName string  `json:"username"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Telephone   *string `json:"telephone"`
}

type registerResponseDto struct {
	common.StandardResponse
}

func newRegisterHandler(accountService services.IAccountService) web.IHandler {
	return &registerHandler{
		accountService: accountService,
	}
}

func (h *registerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &registerHandlerRequestDto{}
	resp := &registerResponseDto{}

	defer func() {
		common.EnableCors(w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}()

	if err := utils.SetRequestDto(r, req); err != nil {
		resp.SetError(err, http.StatusBadRequest)
		return
	}

	_, err := h.accountService.CreateAccount(
		r.Context(),
		req.AccountName,
		req.Email,
		req.Password,
		req.Telephone,
	)
	if err != nil {
		resp.SetError(err, http.StatusBadGateway)
		return
	}
}

func (h *registerHandler) GetMethod() string {
	return http.MethodPost
}

func (h *registerHandler) GetPath() string {
	return root + "/register/"
}
