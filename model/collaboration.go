package model

type Collaboration struct {
	ID         string `gorm:"type:varchar(50);primaryKey"`
	UserId     string `gorm:"type:varchar(50);primaryKey"`
	PlaylistID string `gorm:"type:varchar(50);primaryKey"`
}
