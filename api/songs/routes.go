package songs

import "github.com/gin-gonic/gin"

func (r Register) Router(router *gin.Engine, handler SongHandler) {
	songs := router.Group(r.Uri)
	songs.Use(r.Middleware...)
	{
		songs.GET("/", handler.GetSongsHandler)
	}
}
