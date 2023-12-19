package config

import "github.com/ProxySafe/site-backend/src/modules/db/configurator/postgres"

type Config struct {
	DB         *DBConfig `yaml:"db"`
	SigningKey string    `yaml:"signing_key"`
	TokenTTL   int64     `yaml:"token_ttl"`
}

type DBConfig struct {
	Base *postgres.DbConfig `yaml:"base"`
}
