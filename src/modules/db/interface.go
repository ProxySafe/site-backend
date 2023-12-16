package db

import (
	"context"
	"database/sql"
)

type IReadOnlyDBManager interface {
	ReadDB() ISQLExecutor
}

type IDBManager interface {
	IReadOnlyDBManager
	WriteDB() ISQLExecutor
}

type ICluster interface {
	AddNode(role NodeRole, addr string) error
	RemoveNode(role NodeRole, addr string) error
	Next(role NodeRole) (ISQLExecutor, error)
}

type IConfigurator interface {
	Configure(cluster ICluster) error
	DriverName() string
}

type ExecutorsMap map[string]ISQLExecutor

type ISQLExecutor interface {
	Run(ctx context.Context, dest interface{}, query IToSQL) error
	Exec(ctx context.Context, query IToSQL) (sql.Result, error)
	NamedExec(ctx context.Context, queryTemplate string, sources []any) error
}

type IToSQL interface {
	ToSql() (sqlStr string, args []interface{}, err error)
}
