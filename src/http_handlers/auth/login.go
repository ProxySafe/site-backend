package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/http_handlers/common"
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
	"github.com/ProxySafe/site-backend/src/utils"
)

type loginHandler struct {
	accountService services.IAccountService
	authService    services.IAuthService
}

type loginRequestDto struct {
	entities.Fingerprint
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponseDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Message      string `json:"message"`
	StatusCode   int    `json:"statusCode"`
	AccountId    int    `json:"account_id"`
}

func (l *loginResponseDto) setError(err error, statusCode int) {
	if err != nil {
		l.Message = err.Error()
	}
	l.StatusCode = statusCode
}

func newLoginHandler(accountService services.IAccountService, authService services.IAuthService) web.IHandler {
	return &loginHandler{
		accountService: accountService,
		authService:    authService,
	}
}

func (h *loginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	requestDto := &loginRequestDto{}
	responseDto := &loginResponseDto{}

	defer func() {
		common.EnableCors(w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseDto)
	}()

	if err := utils.SetRequestDto(r, requestDto); err != nil {
		responseDto.setError(err, http.StatusBadRequest)
		return
	}

	if len(requestDto.Username) == 0 || len(requestDto.Password) == 0 {
		// TODO: add mistakes types
		responseDto.setError(fmt.Errorf("invalid username of password"), http.StatusBadRequest)
		return
	}

	account, err := h.accountService.GetByUsername(r.Context(), requestDto.Username)
	if err != nil {
		targetErr := sql.ErrNoRows
		if errors.Is(err, targetErr) {
			responseDto.setError(err, http.StatusNotFound)
			return
		}

		responseDto.setError(err, http.StatusInternalServerError)
		return
	}

	if account.HashedPassword != utils.GetPasswordHash(requestDto.Password) {
		// TODO: add mistakes types
		responseDto.setError(fmt.Errorf("incorrect password"), http.StatusBadRequest)
		return
	}

	accessToken, err := h.authService.GenerateAccessToken(r.Context(), requestDto.Username)
	if err != nil {
		responseDto.setError(err, http.StatusInternalServerError)
		return
	}

	refreshToken, err := h.authService.GenerateRefreshToken(r.Context(), int64(account.Id), requestDto.Fingerprint)
	if err != nil {
		responseDto.setError(err, http.StatusInternalServerError)
		return
	}

	responseDto.AccessToken = accessToken
	responseDto.RefreshToken = refreshToken
	responseDto.AccountId = account.Id
	responseDto.setError(nil, http.StatusOK)
}

func (h *loginHandler) GetMethod() string {
	return http.MethodPost
}

func (h *loginHandler) GetPath() string {
	return root + "/login/"
}
