package core

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/ProxySafe/site-backend/src/modules/db"
	"github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
)

type executor struct {
	db *sqlx.DB
}

func NewExecutor(driverName, dataSourceName string) (db.ISQLExecutor, error) {
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &executor{
		db: db,
	}, nil
}

func (e *executor) Run(
	ctx context.Context,
	dest interface{},
	query db.IToSQL,
) error {
	q, a, err := query.ToSql()
	if err != nil {
		return err
	}

	rows, err := e.db.QueryContext(ctx, q, a...)
	if err != nil {
		return err
	}

	return scan.Rows(dest, rows)
}

func (e *executor) Exec(ctx context.Context, query db.IToSQL) (sql.Result, error) {
	queryString, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	return e.db.Exec(queryString, args...)
}

func (e *executor) NamedExec(ctx context.Context, queryTemplate string, sources []any) error {
	for _, source := range sources {
		_, err := e.db.NamedExecContext(ctx, queryTemplate, source)
		if err != nil {
			return err
		}
	}
	return nil
}
