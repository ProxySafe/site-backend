package db

import "context"

type IReadOnlyDBManager interface {
	ReadDB(ctx context.Context)
}

type IDBManager interface {
}

type IReadOnlySQLExecutor interface {
}
