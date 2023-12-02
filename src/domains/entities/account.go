package entities

import (
	"database/sql"
	"time"
)

type Account struct {
	Id             int            `db:"id"`
	Name           string         `db:"name"`
	HashedPassword string         `db:"hashed_password"`
	Email          string         `db:"email"`
	Telephone      sql.NullString `db:"telephone"`
	CreatedAt      time.Time      `db:"created_at"`
	DeletedAt      sql.NullTime   `db:"deleted_at"`
	Enabled        uint8          `db:"enabled"`
	ChatId         sql.NullString `db:"chat_id"`
}
