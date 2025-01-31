package domain

import (
	"time"
)

type RefreshSession struct {
	ID           int64     `db:"id" json:"id,omitempty"`
	UserID       int64     `db:"user_id" json:"userId,omitempty"`
	RefreshToken string    `db:"refresh_token" json:"refreshToken,omitempty"`
	UserAgent    string    `db:"user_agent" json:"userAgent,omitempty"`
	ExpiresIn    int64     `db:"expires_in" json:"expiresIn,omitempty"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt,omitempty"`
}
