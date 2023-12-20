package repositories

import (
	"context"
	"database/sql"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/modules/db"
	"github.com/elgris/sqrl"
)

const (
	refreshTokenTableName = "refresh_token"
)

type refreshTokenRepository struct {
	db db.IDBManager
}

func NewRefreshTokenRepository(db db.IDBManager) IRefreshTokenRepository {
	return &refreshTokenRepository{
		db: db,
	}
}

func (r *refreshTokenRepository) Add(ctx context.Context, token *entities.RefreshToken) error {
	q := sqrl.Insert(refreshTokenTableName).
		SetMap(token.GetFieldsMap()).
		PlaceholderFormat(sqrl.Dollar)

	ex := r.db.WriteDB()
	if _, err := ex.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *refreshTokenRepository) FindByAccountId(
	ctx context.Context,
	accountId int64,
) (*entities.RefreshToken, error) {
	q := sqrl.Select("*").From(refreshTokenTableName).
		Where(sqrl.Eq{"account_id": accountId}).PlaceholderFormat(sqrl.Dollar)

	dest := &entities.RefreshToken{}
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, dest, q); err != nil {
		return nil, err
	}

	return dest, nil
}

func (r *refreshTokenRepository) FindByUsername(
	ctx context.Context,
	username string,
) (*entities.RefreshToken, error) {
	q := sqrl.Select("*").From(refreshTokenTableName + " r").
		LeftJoin(accountTableName + " a ON a.id = r.account_id").
		Where(sqrl.Eq{"a.name": username}).PlaceholderFormat(sqrl.Dollar)

	var dest []entities.RefreshToken
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, &dest, q); err != nil {
		return nil, err
	}

	if len(dest) == 0 {
		return nil, sql.ErrNoRows
	}
	return &dest[0], nil
}

func (r *refreshTokenRepository) Remove(ctx context.Context, refreshToken *entities.RefreshToken) error {
	q := sqrl.Delete(refreshTokenTableName).
		Where(sqrl.Eq{"account_id": refreshToken.AccountId}).PlaceholderFormat(sqrl.Dollar)

	ex := r.db.WriteDB()
	if _, err := ex.Exec(ctx, q); err != nil {
		return err
	}
	return nil
}

func (r *refreshTokenRepository) RemoveByUsername(ctx context.Context, username string) error {
	q := sqrl.Delete().
		From(refreshTokenTableName + " USING " + refreshTokenTableName + " AS r").
		LeftJoin(accountTableName + " a ON a.id = r.account_id").
		Where(sqrl.Eq{"a.name": username}).PlaceholderFormat(sqrl.Dollar)

	ex := r.db.WriteDB()
	if _, err := ex.Exec(ctx, q); err != nil {
		return err
	}
	return nil
}
