package entities

import "time"

type Message struct {
	Id           int       `db:"id"`
	AccountId    int       `db:"account_id"`
	SendTime     time.Time `db:"send_time"`
	Text         string    `json:"text"`
	ChatId       int       `json:"chat_id"`
	MessageIndex int       `json:"message_index"`
}
