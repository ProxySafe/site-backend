package repositories

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
)

type IAccountRepository interface {
	FindAll(ctx context.Context) ([]entities.Account, error)
	FindByEmail(ctx context.Context, email string) (*entities.Account, error)
	FindByUsername(ctx context.Context, accountName string) (*entities.Account, error)
	Add(ctx context.Context, account *entities.Account) error
}

type ICountryProxyRepository interface {
	FindByProxy(ctx context.Context, proxyId int64) ([]entities.CountryProxy, error)
}

type IMessageRepository interface {
	FindByAccountId(ctx context.Context, accountId int64) ([]entities.Message, error)
}

type IOrderRepository interface {
	FindAll(ctx context.Context) ([]entities.Order, error)
	FindByAccountId(ctx context.Context, accountId int64) ([]entities.Order, error)
	Add(ctx context.Context, order *entities.Order) (int64, error)
}

type IProtocolProxyRepository interface {
	FindByProxy(ctx context.Context, proxyId int64) ([]entities.ProtocolProxy, error)
}

type IProxyRepository interface {
	FindAll(ctx context.Context) ([]entities.Proxy, error)
	FindByAccountId(ctx context.Context, accountId int64) ([]entities.Proxy, error)
	FindNoBusy(ctx context.Context) ([]entities.Proxy, error)
	SetToOrder(ctx context.Context, orderId int, proxies []entities.Proxy) error
}

type ITCpProxyRepository interface {
	FindByProxy(ctx context.Context, proxyId int64) ([]entities.TCpProxy, error)
}

type IRefreshTokenRepository interface {
	Add(ctx context.Context, refreshToken *entities.RefreshToken) error
	FindByAccountId(ctx context.Context, accountId int64) (*entities.RefreshToken, error)
	FindByUsername(ctx context.Context, username string) (*entities.RefreshToken, error)
	Remove(ctx context.Context, refreshToken *entities.RefreshToken) error
	RemoveByUsername(ctx context.Context, username string) error
}
