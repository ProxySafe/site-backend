package repositories

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
)

type IAccountRepository interface {
	FindAll(ctx context.Context) ([]entities.Account, error)
	FindByEmail(ctx context.Context, email string) (*entities.Account, error)
	FindByUsername(ctx context.Context, accountName string) (*entities.Account, error)
}

type ICountryProxyRepository interface {
	FindByProxy(ctx context.Context, proxyId int64) ([]entities.CountryProxy, error)
}

type IMessageRepository interface {
	FindByAccountId(ctx context.Context, accountId int64) ([]entities.Message, error)
}

type IOrderRepository interface {
	FindByAccountId(ctx context.Context, accountId int64) ([]entities.Order, error)
}

type IProtocolProxyRepository interface {
	FindByProxy(ctx context.Context, proxyId int64) ([]entities.ProtocolProxy, error)
}

type IProxyRepository interface {
	FindAll(ctx context.Context) ([]entities.Proxy, error)
	FindByProvider(ctx context.Context, providerName string) ([]entities.Proxy, error)
	FindByAccountId(ctx context.Context, accountId int64) ([]entities.Proxy, error)
}

type ITCpProxyRepository interface {
	FindByProxy(ctx context.Context, proxyId int64) ([]entities.TCpProxy, error)
}
