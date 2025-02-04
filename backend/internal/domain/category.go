package domain

type Category struct {
	Id   int64  `db:"id" json:"id,omitempty"`
	Name string `db:"name" json:"name,omitempty"`
}
