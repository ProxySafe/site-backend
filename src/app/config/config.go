package config

import "github.com/ProxySafe/site-backend/src/modules/db/configurator/postgres"

type Config struct {
	DB *DBConfig `yaml:"db"`
}

type DBConfig struct {
	Base *postgres.DbConfig `yaml:"base"`
}
