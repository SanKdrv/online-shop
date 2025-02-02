package domain

import "time"

type ProductImage struct {
	ID        int64     `db:"id" json:"id,omitempty"`
	ProductID int64     `db:"product_id" json:"product_id,omitempty"`
	ImageHash string    `db:"image_hash" json:"image_hash,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
