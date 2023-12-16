package home

import "github.com/ProxySafe/site-backend/src/modules/web"

func NewHandlers() []web.IHandler {
	return []web.IHandler{
		newHomeHandler(),
	}
}
