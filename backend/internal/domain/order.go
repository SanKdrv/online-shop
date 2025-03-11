package domain

import "time"

type Order struct {
	Id        int64     `db:"id" json:"id,omitempty"`
	UserId    int64     `db:"user_id" json:"user_id,omitempty"`
	Status    string    `db:"status" json:"status,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
