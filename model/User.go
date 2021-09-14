package model

type User struct {
	ID           string     `gorm:"type:varchar(50);primaryKey"`
	Username     string     `gorm:"type:varchar(250);unique" binding:"required"`
	Password     string     `gorm:"type:text" binding:"required"`
	FullName     string     `gorm:"type:text" binding:"required"`
	Playlist     *Playlist  `gorm:"foreignKey:owner;reference:ID"`
	Playlists    []Playlist `gorm:"many2many:collaborations"`
	RefreshToken RefreshToken
}
