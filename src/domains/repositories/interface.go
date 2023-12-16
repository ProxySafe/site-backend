package repositories

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
)

type IAccountRepository interface {
	GetAll(ctx context.Context) ([]entities.Account, error)
	GetAccount(ctx context.Context, email string) (*entities.Account, error)
}

type ICountryProxyRepository interface {
}

type IMessageRepository interface {
}

type IOrderRepository interface {
}

type IProtocolProxyRepository interface {
}

type IProxyRepository interface {
}

type ITCpProxyRepository interface {
}
