package entities

import (
	"time"
)

type Account struct {
	Id             int       `db:"id" json:"id"`
	Name           string    `db:"name" json:"name"`
	HashedPassword string    `db:"hashed_password" json:"hashed_password"`
	Email          string    `db:"email" json:"email"`
	Telephone      *string   `db:"telephone" json:"telephone"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	DeletedAt      *string   `db:"deleted_at" json:"deleted_at"`
	Enabled        bool      `db:"enabled" json:"enabled"`
	ChatId         *string   `db:"chat_id" json:"chat_id"`
}

func NewAccountService(accountName, hashedPassword, email string, telephone *string) *Account {
	return &Account{
		Name:           accountName,
		HashedPassword: hashedPassword,
		Email:          email,
		Telephone:      telephone,
		CreatedAt:      time.Now(),
	}
}

func (a *Account) GetFieldsMap() map[string]interface{} {
	return map[string]interface{}{
		"name":            a.Name,
		"hashed_password": a.HashedPassword,
		"email":           a.Email,
		"enabled":         a.Enabled,
		"created_at":      a.CreatedAt,
	}
}
