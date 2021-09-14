package migration

import (
	"kurnia123/restfull-api-v1/app"
	"kurnia123/restfull-api-v1/helper"
	"kurnia123/restfull-api-v1/model"
	"testing"

	"github.com/joho/godotenv"
)

func TestMigrateDatabase(t *testing.T) {
	godotenv.Load("../.env")
	db := app.NewGorm(app.NewDB())

	mysqlDB, err := db.DB()
	helper.PanicError(err)

	defer mysqlDB.Close()

	err = db.AutoMigrate(
		&model.Song{},
		&model.User{},
		&model.Playlist{},
		&model.PlaylistSong{},
		&model.Collaboration{},
		&model.RefreshToken{},
	)
	helper.PanicError(err)

	err = db.SetupJoinTable(&model.User{}, "Playlists", &model.Collaboration{})
	helper.PanicError(err)

	err = db.SetupJoinTable(&model.Playlist{}, "Songs", &model.PlaylistSong{})
	helper.PanicError(err)
}
