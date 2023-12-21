package services

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
)

type IAccountService interface {
	GetAll(ctx context.Context) ([]entities.Account, error)
	GetByUsername(ctx context.Context, userName string) (*entities.Account, error)
	CreateAccount(ctx context.Context, userName, email, password string, telephone *string) (*entities.Account, error)
}

type IAuthService interface {
	GenerateAccessToken(ctx context.Context, userName string) (string, error)
	GenerateRefreshToken(
		ctx context.Context,
		accountId int64,
		fingerprint entities.Fingerprint,
	) (string, error)
	ParseToken(ctx context.Context, accessToken string) (string, bool, error)
	RefreshAccessToken(
		ctx context.Context,
		oldAccessToken,
		refreshToken string,
		fingerprint entities.Fingerprint,
	) (string, error)
	RemoveRefreshToken(ctx context.Context, accessToken string, fp *entities.Fingerprint) error
}

type IProxyService interface {
	GetAll(ctx context.Context) ([]entities.Proxy, error)
	GetByAccount(ctx context.Context, accountId int64) ([]entities.Proxy, error)
	GetNoBusy(ctx context.Context) ([]entities.Proxy, error)
	GetProxiesByAmount(ctx context.Context, amount int) ([]entities.Proxy, error)
}

type IOrderService interface {
	GetAll(ctx context.Context) ([]entities.Order, error)
	GetByAccount(ctx context.Context, accountId int64) ([]entities.Order, error)
	CreateOrderByProxies(ctx context.Context, period, accountId int, proxies []entities.Proxy) error
}

// TODO: for 2-factor authorization
type IEmailService interface {
}
