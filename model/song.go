package model

type Song struct {
	ID        string `gorm:"type:varchar(50);primaryKey"`
	Title     string `gorm:"type:text"`
	Year      uint
	Performer string
	Genre     string
	Duration  uint
}
