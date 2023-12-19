package entities

import (
	"time"
)

type Proxy struct {
	Id             int        `db:"id" json:"id"`
	Addr           string     `db:"addr" json:"addr"`
	Enabled        uint8      `db:"enabled" json:"enabled"`
	TcpFingerprint string     `db:"tcp_fingerprint" json:"tcp_fingerprint"`
	Country        string     `db:"country" json:"country"`
	ServerName     string     `db:"server_name" json:"server_name"`
	IsBusy         uint8      `db:"is_busy" json:"is_busy"`
	OrderId        *int64     `db:"order_id" json:"order_id"`
	Speed          *int64     `db:"speed" json:"speed"`
	ExternalIP     string     `db:"external_ip" json:"external_ip"`
	User           string     `db:"user" json:"user"`
	Password       string     `db:"password" json:"password"`
	RotationPeriod uint8      `db:"rotation_period" json:"rotation_period"`
	RentedAt       *time.Time `db:"rented_at" json:"rented_at"`
	RentFinish     *time.Time `db:"rent_finish" json:"rent_finish"`
	Protocol       string     `db:"protocol" json:"protocol"`
	Price          int        `db:"price" json:"price"`
}

type IdWithProxyId struct {
	Id      int `db:"id" json:"id"`
	ProxyId int `db:"proxy_id" json:"proxy_id"`
}

type TCpProxy struct {
	IdWithProxyId
	Fingerprint string `db:"fingerprint" json:"fingerprint"`
}

type ProtocolProxy struct {
	IdWithProxyId
	Protocol string `db:"protocol" json:"protocol"`
}

type CountryProxy struct {
	IdWithProxyId
	Country string `db:"country" json:"country"`
}
