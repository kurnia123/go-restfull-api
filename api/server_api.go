package api

import (
	"github.com/gin-gonic/gin"
)

type ServerAPI struct {
	Router *gin.Engine
}

func (s *ServerAPI) Register(api ...APIRegister) {
	for _, handler := range api {
		handler.Load(s.Router)
	}
}
