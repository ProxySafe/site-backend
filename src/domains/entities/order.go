package entities

import "time"

type Order struct {
	Id                  int       `db:"id" json:"id"`
	OrderDate           time.Time `db:"order_date" json:"order_date"`
	AccountId           int       `db:"account_id" json:"account_id"`
	OrderExpirationDate time.Time `db:"order_expiration_date" json:"order_expiration_date"`
	ProxiesNumber       int       `db:"proxies_number" json:"proxies_number"`
}

func (o *Order) GetFieldsMap() map[string]interface{} {
	return map[string]interface{}{
		"order_date":            o.OrderDate,
		"account_id":            o.AccountId,
		"order_expiration_date": o.OrderExpirationDate,
		"proxies_number":        o.ProxiesNumber,
	}
}
