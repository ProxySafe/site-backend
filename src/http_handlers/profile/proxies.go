package profile

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/http_handlers/common"
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
	"github.com/ProxySafe/site-backend/src/utils"
)

type proxiesHandler struct {
	proxyService services.IProxyService
}

type proxiesRequestDto struct {
	AccountId int64 `json:"account_id"`
}

type proxiesResponseDto struct {
	common.StandardResponse
	Proxies     []entities.Proxy `json:"proxies"`
	AccessToken string           `json:"accessToken"`
}

func newProxiesHandler(proxyService services.IProxyService) web.IHandler {
	return &proxiesHandler{
		proxyService: proxyService,
	}
}

func (h *proxiesHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &proxiesRequestDto{}
	resp := &proxiesResponseDto{}

	defer func() {
		common.EnableCors(w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}()

	if err := utils.SetRequestDto(r, req); err != nil {
		resp.SetError(err, http.StatusBadRequest)
		return
	}

	proxies, err := h.proxyService.GetByAccount(r.Context(), req.AccountId)
	if !errors.Is(err, sql.ErrNoRows) {
		resp.SetError(err, http.StatusBadGateway)
		return
	}

	resp.Proxies = proxies
	resp.SetError(nil, http.StatusOK)
}

func (h *proxiesHandler) GetMethod() string {
	return http.MethodPost
}

func (h *proxiesHandler) GetPath() string {
	return root + "/proxies/"
}
