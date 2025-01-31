package domain

import (
	"time"
)

type User struct {
	ID           int64     `db:"id" json:"id,omitempty"`
	Username     string    `db:"username" json:"username,omitempty"`
	Name         string    `db:"name" json:"name,omitempty"`
	GoogleName   string    `db:"google_name" json:"googleName,omitempty"`
	Email        string    `db:"email" json:"email,omitempty"`
	PasswordHash string    `db:"password_hash" json:"passwordHash,omitempty"`
	Role         string    `db:"role" json:"role,omitempty" gorm:"default:user"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt,omitempty"`
	UpdatedAt    time.Time `db:"updated_at" json:"updatedAt,omitempty"`
}
