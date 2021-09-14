package users

import (
	"fmt"
	"kurnia123/restfull-api-v1/schema"
	"kurnia123/restfull-api-v1/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		Service: userService,
	}
}

func (u *UserHandler) PostUserHandler(c *gin.Context) {
	var sigin schema.User

	err := c.Bind(&sigin)
	fmt.Println("DATA : ", err)
	fmt.Println("DATA : ", sigin)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	u.Service.Create(c, &sigin)
}

func (u *UserHandler) GetUserByIdHandler(c *gin.Context) {
	id := c.Param("id")
	u.Service.FindById(c, id)
}

func (u *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	u.Service.FindByUsername(c, username)
}

func (u *UserHandler) GetUserAll(c *gin.Context) {
	u.Service.FindAll(c)
}
