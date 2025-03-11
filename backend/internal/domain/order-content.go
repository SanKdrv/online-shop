package domain

type OrderContent struct {
	Id        int64   `db:"id" json:"id,omitempty"`
	OrderId   int64   `db:"order_id" json:"order_id,omitempty"`
	ProductId int64   `db:"product_id" json:"product_id,omitempty"`
	Count     int64   `db:"count" json:"count,omitempty"`
	SumPrice  float64 `db:"sum_price" json:"sum_price,omitempty"`
}
