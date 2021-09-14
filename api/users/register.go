package users

import (
	"github.com/gin-gonic/gin"
)

type Register struct {
	Uri        string
	Middleware []gin.HandlerFunc
	Handler    UserHandler
}

func (r Register) Load(router *gin.Engine) {
	r.Router(router, r.Handler)
}
