package domain

import "time"

type Product struct {
	ID          int       `db:"id" json:"id"`
	BrandID     int       `db:"brand_id" json:"brand_id"`
	CategoryID  int       `db:"category_id" json:"category_id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Price       float64   `db:"price" json:"price"`
	CreatedAt   time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
