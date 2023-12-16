package repositories

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/modules/db"
	"github.com/elgris/sqrl"
)

const (
	accountTableName = "account"
)

type accountRepository struct {
	db db.IDBManager
}

func NewAccountRepository(db db.IDBManager) IAccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) GetAll(ctx context.Context) ([]entities.Account, error) {
	q := sqrl.Select("id, name, hashed_password").From(accountTableName)

	var dest []entities.Account
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, &dest, q); err != nil {
		return nil, err
	}
	return dest, nil
}

func (r *accountRepository) GetAccount(ctx context.Context, email string) (*entities.Account, error) {

	return nil, nil
}
