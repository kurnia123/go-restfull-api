package model

type Playlist struct {
	ID    string `gorm:"type:varchar(50);primaryKey"`
	Name  string `gorm:"type:text"`
	Owner string `gorm:"type:varchar(50)"`
	Songs []Song `gorm:"many2many:playlist_songs"`
}
