package main

import (
	"kurnia123/restfull-api-v1/api"
	"kurnia123/restfull-api-v1/api/authentications"
	"kurnia123/restfull-api-v1/api/songs"
	"kurnia123/restfull-api-v1/api/users"
	"kurnia123/restfull-api-v1/app"
	"kurnia123/restfull-api-v1/exeption"
	"kurnia123/restfull-api-v1/middleware"
	"kurnia123/restfull-api-v1/repository"
	"kurnia123/restfull-api-v1/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	router := gin.Default()
	validate := validator.New()
	db := app.NewGorm(app.NewDB())

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userHandler := users.NewUserHandler(userService)

	songRepository := repository.NewSongRepository()
	songService := service.NewSongService(songRepository, db, validate)
	songHandler := songs.NewSongHandler(songService)

	authenticationRepository := repository.NewAuthRepository()
	authenticationService := service.NewAuthService(authenticationRepository, db, validate)
	authenticationHandler := authentications.NewAuthHandler(authenticationService)

	server := api.ServerAPI{
		Router: router,
	}
	// Set All Route use Middleware
	// router.Use(gin.CustomRecovery(exeption.ErrorHandler))
	router.Use(exeption.ErrorHandler)

	server.Register(
		users.Register{
			Uri: "/users",
			Middleware: []gin.HandlerFunc{
				middleware.AuthMiddleware,
			},
			Handler: userHandler,
		},
		songs.Register{
			Uri: "/songs",
			Middleware: []gin.HandlerFunc{
				middleware.AuthMiddleware,
			},
			Handler: songHandler,
		},
		authentications.Register{
			Uri:     "/authentications",
			Handler: authenticationHandler,
		},
	)

	router.Run()
}
