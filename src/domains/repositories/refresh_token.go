package repositories

import (
	"context"

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
	q := sqrl.Insert(refreshTokenTableName).SetMap(token.GetFieldsMap())

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
		Where(sqrl.Eq{"account_id": accountId})

	dest := &entities.RefreshToken{}
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, dest, q); err != nil {
		return nil, err
	}

	return dest, nil
}
