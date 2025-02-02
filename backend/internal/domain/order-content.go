package domain

type OrderContent struct {
	ID        int64   `db:"id" json:"id,omitempty"`
	OrderID   int64   `db:"order_id" json:"order_id,omitempty"`
	ProductID int64   `db:"product_id" json:"product_id,omitempty"`
	Count     float64 `db:"count" json:"count,omitempty"`
}
