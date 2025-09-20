package repositories

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/modules/db"
	"github.com/elgris/sqrl"
	"github.com/elgris/sqrl/pg"
)

const (
	proxyTableName = "proxy"
)

type proxyRepository struct {
	db db.IDBManager
}

func NewProxyRepository(db db.IDBManager) IProxyRepository {
	return &proxyRepository{
		db: db,
	}
}

func (r *proxyRepository) FindAll(ctx context.Context) ([]entities.Proxy, error) {
	q := sqrl.Select("*").From(proxyTableName)

	var dest []entities.Proxy
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, &dest, q); err != nil {
		return nil, err
	}

	return dest, nil
}

func (r *proxyRepository) FindByAccountId(ctx context.Context, accountId int64) ([]entities.Proxy, error) {
	q := sqrl.Select("p.*").
		From(proxyTableName + " p").
		LeftJoin(orderTableName + " o ON o.id = p.order_id").
		Where(sqrl.Eq{"o.account_id": accountId}).
		PlaceholderFormat(sqrl.Dollar)

	var dest []entities.Proxy
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, &dest, q); err != nil {
		return nil, err
	}

	return dest, nil
}

func (r *proxyRepository) FindNoBusy(ctx context.Context) ([]entities.Proxy, error) {
	q := sqrl.Select("*").
		From(proxyTableName).
		Where(sqrl.Eq{
			"is_busy": 0,
		}).
		PlaceholderFormat(sqrl.Dollar)

	var dest []entities.Proxy
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, &dest, q); err != nil {
		return nil, err
	}

	return dest, nil
}

func (r *proxyRepository) SetToOrder(
	ctx context.Context,
	orderId int,
	proxies []entities.Proxy,
) error {
	ids := entities.Proxies(proxies).GetIds()

	q := sqrl.Update(proxyTableName).
		Set("is_busy", 1).
		Set("order_id", orderId).
		Where(sqrl.Expr("id = any($1)", pg.Array(ids))).
		PlaceholderFormat(sqrl.Dollar)

	ex := r.db.WriteDB()
	if _, err := ex.Exec(ctx, q); err != nil {
		return err
	}
	return nil
}
