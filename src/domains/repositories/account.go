package repositories

import (
	"context"
	"database/sql"

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
	q := sqrl.Select("*").
		From(accountTableName).
		Where(sqrl.Eq{"email": email}).
		PlaceholderFormat(sqrl.Dollar)

	dest := &entities.Account{}
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, dest, q); err != nil {
		return nil, err
	}
	return dest, nil
}

func (r *accountRepository) FindByUsername(ctx context.Context, userName string) (*entities.Account, error) {
	q := sqrl.Select("*").
		From(accountTableName).
		Where(sqrl.Eq{"name": userName}).
		PlaceholderFormat(sqrl.Dollar)

	var dest []entities.Account
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, &dest, q); err != nil {
		return nil, err
	}

	if len(dest) > 1 {
		return nil, &ErrManyUsers{}
	}

	if len(dest) == 0 {
		return nil, sql.ErrNoRows
	}
	return &dest[0], nil
}

func (r *accountRepository) Add(ctx context.Context, account *entities.Account) error {
	q := sqrl.Insert(accountTableName).
		SetMap(account.GetFieldsMap()).
		PlaceholderFormat(sqrl.Dollar)

	ex := r.db.ReadDB()
	if _, err := ex.Exec(ctx, q); err != nil {
		return err
	}
	return nil
}
