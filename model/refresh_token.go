package model

type RefreshToken struct {
	ID     string `gorm:"type:varchar(50);primaryKey"`
	Token  string `gorm:"type:varchar(250)"`
	UserID string
}
