package app

import (
	"net/http"
	"strconv"

	"github.com/ProxySafe/site-backend/src/app/config"
	"github.com/ProxySafe/site-backend/src/app/repositories"
	"github.com/ProxySafe/site-backend/src/app/resources"
	"github.com/ProxySafe/site-backend/src/app/services"
	"github.com/ProxySafe/site-backend/src/http_handlers/auth"
	"github.com/ProxySafe/site-backend/src/modules/web"
	"github.com/gorilla/mux"
)

type App struct {
	cfg          *config.Config
	resources    *resources.Resources
	repositories *repositories.Repositories
	services     *services.Services
	webServer    *mux.Router
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

	handlers = append(handlers, auth.NewHandlers(a.services.AccountService, a.services.AuthService)...)
	return handlers
}

func (a *App) Init() {
	a.resources = resources.NewResources(a.cfg)
	a.repositories = repositories.NewRepositories(a.resources)
	a.services = services.NewServices(a.resources, a.repositories)
	a.initWebServer(a.initHttpHandlers())
}

func (a *App) Run(port int) {
	http.ListenAndServe(":"+strconv.Itoa(port), a.webServer)
}
