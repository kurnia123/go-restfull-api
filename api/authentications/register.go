package authentications

import "github.com/gin-gonic/gin"

type Register struct {
	Uri        string
	Middleware []gin.HandlerFunc
	Handler    AuthHandler
}

func (r Register) Load(router *gin.Engine) {
	r.Router(router, r.Handler)
}
