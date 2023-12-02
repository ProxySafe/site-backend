package entities

import "time"

type Order struct {
	Id                  int       `db:"id"`
	OrderDate           time.Time `db:"order_date"`
	AccountId           int       `db:"account_id"`
	OrderExpirationDate time.Time `db:"order_expiration_date"`
	ProxiesNumber       int       `db:"proxies_number"`
}