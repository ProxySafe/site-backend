package services

import (
	"github.com/ProxySafe/site-backend/src/app/repositories"
	"github.com/ProxySafe/site-backend/src/app/resources"
	"github.com/ProxySafe/site-backend/src/services"
	"github.com/ProxySafe/site-backend/src/services/account"
	"github.com/ProxySafe/site-backend/src/services/auth"
)

type Services struct {
	AccountService services.IAccountService
	AuthService    services.IAuthService
}

func NewServices(res *resources.Resources, repos *repositories.Repositories) *Services {
	return &Services{
		AccountService: account.NewService(repos.AccountRepository),
		AuthService:    auth.NewService(res.SigningKey, res.TokenTTL, repos.RefreshTokenRepository),
	}
}
