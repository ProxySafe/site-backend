package repositories

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/modules/db"
)

type accountRepository struct {
	manager db.IDBManager
}

func NewAccountRepository(manager db.IDBManager) IAccountRepository {
	return &accountRepository{
		manager: manager,
	}
}

func (r *accountRepository) GetAccount(ctx context.Context, email string) (*entities.Account, error) {

	return nil, nil
}
