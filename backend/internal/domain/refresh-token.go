package domain

type RefreshToken struct {
	RefreshToken string `json:"refresh-token" db:"refresh_token" gorm:"column:refresh_token"`
}
