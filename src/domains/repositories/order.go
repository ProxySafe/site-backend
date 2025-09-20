package repositories

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/modules/db"
	"github.com/elgris/sqrl"
)

const (
	orderTableName = "proxy_order"
)

type orderRepository struct {
	db db.IDBManager
}

func NewOrderRepository(db db.IDBManager) IOrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) FindAll(ctx context.Context) ([]entities.Order, error) {
	q := sqrl.Select("*").From(orderTableName)

	var dest []entities.Order
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, &dest, q); err != nil {
		return nil, err
	}

	return dest, nil
}

func (r *orderRepository) FindByAccountId(ctx context.Context, accountId int64) ([]entities.Order, error) {
	q := sqrl.Select("*").
		From(orderTableName).
		Where(sqrl.Eq{"account_id": accountId}).
		PlaceholderFormat(sqrl.Dollar)

	var dest []entities.Order
	ex := r.db.ReadDB()
	if err := ex.Run(ctx, &dest, q); err != nil {
		return nil, err
	}

	return dest, nil
}

func (r *orderRepository) Add(ctx context.Context, order *entities.Order) (int64, error) {
	q := sqrl.Insert(orderTableName).
		SetMap(order.GetFieldsMap()).
		PlaceholderFormat(sqrl.Dollar).
		Returning("id")

	ex := r.db.WriteDB()
	res, err := ex.Exec(ctx, q)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}
