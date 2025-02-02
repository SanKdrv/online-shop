package domain

import "time"

type Order struct {
	ID        int64     `db:"id" json:"id,omitempty"`
	OrderID   int64     `db:"order_id" json:"order_id,omitempty"`
	Cost      float64   `db:"cost" json:"cost,omitempty"`
	Status    string    `db:"status" json:"status,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
