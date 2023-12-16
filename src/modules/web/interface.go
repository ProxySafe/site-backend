package web

import (
	"net/http"
)

type IHandler interface {
	GetPath() string
	GetMethod() string
	Handle(w http.ResponseWriter, r *http.Request)
}
