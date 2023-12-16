package services

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
)

type IAccountService interface {
	GetAll(ctx context.Context) ([]entities.Account, error)
}
