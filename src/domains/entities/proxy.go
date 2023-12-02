package entities

import "database/sql"

type Proxy struct {
	Id             int           `db:"id"`
	Addr           string        `db:"addr"`
	Enabled        uint8         `db:"enabled"`
	TcpFingerprint string        `db:"tcp_fingerprint"`
	Country        string        `db:"country"`
	ServerName     string        `db:"server_name"`
	IsBusy         uint8         `db:"is_busy"`
	OrderId        sql.NullInt64 `db:"order_id"`
	Speed          sql.NullInt64 `db:"speed"`
	ExternalIP     string        `db:"external_ip"`
	User           string        `db:"user"`
	Password       string        `db:"password"`
	RotationPeriod uint8         `db:"rotation_period"`
	RentedAt       sql.NullTime  `db:"rented_at"`
	RentFinish     sql.NullTime  `db:"rent_finish"`
	Protocol       string        `db:"protocol"`
	Price          int           `db:"price"`
}

type IdWithProxyId struct {
	Id      int `db:"id"`
	ProxyId int `db:"proxy_id"`
}

type TCpProxy struct {
	IdWithProxyId
	Fingerprint string `db:"fingerprint"`
}

type ProtocolProxy struct {
	IdWithProxyId
	Protocol string `db:"protocol"`
}

type CountryProxy struct {
	IdWithProxyId
	Country string `db:"country"`
}
