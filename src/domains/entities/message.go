package entities

import "time"

type Message struct {
	Id           int       `db:"id" json:"id"`
	AccountId    int       `db:"account_id" json:"account_id"`
	SendTime     time.Time `db:"send_time" json:"send_time"`
	Text         string    `json:"text" json:"text"`
	ChatId       int       `json:"chat_id" json:"chat_id"`
	MessageIndex int       `json:"message_index" json:"message_index"`
}
