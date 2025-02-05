package domain

import "time"

type WebSocket struct {
	Id          int64     `db:"id" json:"Id,omitempty"`
	ChatOwnerId int64     `db:"chat_owner_id" json:"chatOwnerId"`
	UserId      int64     `db:"user_id" json:"userId"`
	Message     string    `db:"message" json:"message"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt,omitempty"`
}
