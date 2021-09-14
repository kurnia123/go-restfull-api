package authentications

import "github.com/gin-gonic/gin"

func Data(d ...string) {

}

func (r Register) Router(router *gin.Engine, handler AuthHandler) {
	songs := router.Group(r.Uri)
	songs.Use(r.Middleware...)
	{
		songs.POST("/", handler.PostAuthHandler)
	}
}
