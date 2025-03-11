package domain

import "time"

type CartContent struct {
	Id        int64     `db:"id" json:"id,omitempty"`
	UserId    int64     `db:"user_id" json:"user_id,omitempty"`
	ProductId int64     `db:"product_id" json:"product_id,omitempty"`
	Count     int       `db:"count" json:"count,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

func (CartContent) TableName() string {
	return "carts_content"
}
