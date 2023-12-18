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

func (r *accountRepository) FindAll(ctx context.Context) ([]entities.Account, error) {
	q := sqrl.Select("*").From(accountTableName)

	var dest []entities.Account
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, &dest, q); err != nil {
		return nil, err
	}
	return dest, nil
}

func (r *accountRepository) FindByEmail(ctx context.Context, email string) (*entities.Account, error) {
	q := sqrl.Select("*").From(accountTableName).Where(sqrl.Eq{"email": email})

	dest := &entities.Account{}
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, dest, q); err != nil {
		return nil, err
	}
	return dest, nil
}

func (r *accountRepository) FindByUsername(ctx context.Context, userName string) (*entities.Account, error) {
	q := sqrl.Select("*").From(accountTableName).Where(sqrl.Eq{"name": userName})

	dest := &entities.Account{}
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, dest, q); err != nil {
		return nil, err
	}
	return dest, nil
}

func (r *accountRepository) Add(ctx context.Context, account *entities.Account) error {
	q := sqrl.Insert(accountTableName).SetMap(account.GetFieldsMap())

	ex := r.db.ReadDB()
	if _, err := ex.Exec(ctx, q); err != nil {
		return err
	}
	return nil
}
