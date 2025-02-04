package domain

import "time"

type Order struct {
	Id        int64     `db:"id" json:"id,omitempty"`
	OrderId   int64     `db:"order_id" json:"order_id,omitempty"`
	Cost      float64   `db:"cost" json:"cost,omitempty"`
	Status    string    `db:"status" json:"status,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
