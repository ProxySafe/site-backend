package store

import (
	"encoding/json"
	"net/http"

	"github.com/ProxySafe/site-backend/src/http_handlers/common"
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/ProxySafe/site-backend/src/services"
	"github.com/ProxySafe/site-backend/src/services/proxy"
	"github.com/ProxySafe/site-backend/src/utils"
)

type buyHandler struct {
	proxyService services.IProxyService
	orderService services.IOrderService
}

type buyRequestDto struct {
	ProxiesAmount int `json:"amount"`
	Period        int `json:"period"`
	AccountId     int `json:"account_id"`
}

type buyResponseDto struct {
	common.StandardResponse
	ProxiesAddrs []string `json:"addrs"`
}

func newBuyHandler(proxyService services.IProxyService, orderService services.IOrderService) web.IHandler {
	return &buyHandler{
		proxyService: proxyService,
		orderService: orderService,
	}
}

func (h *buyHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := &buyRequestDto{}
	resp := &buyResponseDto{}

	defer func() {
		common.EnableCors(w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}()

	if err := utils.SetRequestDto(r, req); err != nil {
		resp.SetError(err, http.StatusBadRequest)
		return
	}

	proxies, err := h.proxyService.GetProxiesByAmount(r.Context(), req.ProxiesAmount)
	if err != nil {
		if proxy.IsErrNoAvailableProxies(err) {
			resp.SetError(err, http.StatusAccepted)
			return
		}

		resp.SetError(err, http.StatusNotFound)
		return
	}

	err = h.orderService.CreateOrderByProxies(r.Context(), req.Period, req.AccountId, proxies)
	if err != nil {
		resp.SetError(err, http.StatusBadGateway)
		return
	}

	resp.SetError(nil, http.StatusOK)
}

func (h *buyHandler) GetMethod() string {
	return http.MethodPost
}

func (h *buyHandler) GetPath() string {
	return root + "/buy/"
}
