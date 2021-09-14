package model

type PlaylistSong struct {
	ID         string `gorm:"type:varchar(50);primaryKey"`
	PlaylistID string `gorm:"type:varchar(50);primaryKey"`
	SongID     string `gorm:"type:varchar(50);primaryKey"`
}
