package types

import "time"

type Request struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Name       string    `json:"name"`
	GoogleName string    `json:"googleName,omitempty"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`

	UserAgent string `json:"userAgent"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password"`
	UserAgent string `json:"userAgent,omitempty"`
}

type SignOutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokensRequest struct {
	RefreshToken string `json:"refresh_token"`
	UserAgent    string `json:"userAgent,omitempty"`
}

type Response struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Status       string `json:"status,omitempty"`
	Error        string `json:"error,omitempty"`
}

type GetUsernameByIDRequest struct {
	UserID int64 `json:"user_id"`
}

type GetUsernameByIDResponse struct {
	Username string `json:"username"`
}
