package songs

import (
	"kurnia123/restfull-api-v1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	SongService service.SongService
}

func NewSongHandler(songService service.SongService) SongHandler {
	return SongHandler{
		SongService: songService,
	}
}

func (s *SongHandler) GetSongsHandler(c *gin.Context) {
	c.String(http.StatusOK, "Ok ini bisa")
}
