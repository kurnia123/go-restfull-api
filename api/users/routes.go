package users

import (
	"github.com/gin-gonic/gin"
)

func (r Register) Router(router *gin.Engine, handler UserHandler) {
	users := router.Group(r.Uri)
	users.Use(r.Middleware...)
	{
		users.GET("/id/:id", handler.GetUserByIdHandler)
		users.GET("/usr/:username", handler.GetUserByUsername)
		users.GET("/all", handler.GetUserAll)
	}

	sigin := router.Group(r.Uri + "/signin")
	{
		sigin.POST("/", handler.PostUserHandler)
	}
}
