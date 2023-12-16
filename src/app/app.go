package app

import (
	"net/http"
	"strconv"

	"github.com/ProxySafe/site-backend/src/app/config"
	"github.com/ProxySafe/site-backend/src/app/resources"
	"github.com/ProxySafe/site-backend/src/http_handlers/home"
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/gorilla/mux"
)

type App struct {
	cfg       *config.Config
	resources *resources.Resources
	webServer *mux.Router
}

func NewApp(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) initWebServer(httpHandlers []web.IHandler) {
	a.webServer = &mux.Router{}
	for _, handler := range httpHandlers {
		a.webServer.HandleFunc(handler.GetPath(), handler.Handle)
	}
}

func (a *App) initHttpHandlers() []web.IHandler {
	var handlers []web.IHandler

	handlers = append(handlers, home.NewHandlers()...)
	return handlers
}

func (a *App) Init() {
	a.resources = resources.NewResources(a.cfg)
	a.initWebServer(a.initHttpHandlers())
}

func (a *App) Run(port int) {
	http.ListenAndServe(":"+strconv.Itoa(port), a.webServer)
}
