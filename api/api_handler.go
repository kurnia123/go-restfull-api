package api

import "github.com/gin-gonic/gin"

type APIRegister interface {
	Load(router *gin.Engine)
}
