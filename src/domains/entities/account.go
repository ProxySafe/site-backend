package entities

import (
	"database/sql"
	"time"
)

type Account struct {
	Id             int            `db:"id" json:"id"`
	Name           string         `db:"name" json:"name"`
	HashedPassword string         `db:"hashed_password" json:"hashed_password"`
	Email          string         `db:"email" json:"email"`
	Telephone      sql.NullString `db:"telephone" json:"telephone"`
	CreatedAt      time.Time      `db:"created_at" json:"created_at"`
	DeletedAt      sql.NullTime   `db:"deleted_at" json:"deleted_at"`
	Enabled        bool           `db:"enabled" json:"enabled"`
	ChatId         sql.NullString `db:"chat_id" json:"chat_id"`
}
