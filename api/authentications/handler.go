package authentications

import (
	"kurnia123/restfull-api-v1/schema"
	"kurnia123/restfull-api-v1/service"
	"kurnia123/restfull-api-v1/tokenize"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service service.AuthService
}

func NewAuthHandler(auth service.AuthService) AuthHandler {
	return AuthHandler{
		Service: auth,
	}
}

func (a *AuthHandler) PostAuthHandler(c *gin.Context) {
	var user schema.User
	var err error
	var accessToken string
	var refreshToken string

	err = c.Bind(&user)
	if err != nil {
		c.Error(err)
		return
	}

	userId := a.Service.AuthenticationUser(c, user.Username, user.Password)
	if userId == "" {
		return
	}

	accessToken, err = tokenize.GenerateAccessToken(user.Username)
	if err != nil {
		c.Error(err)
		return
	}
	refreshToken, err = tokenize.GenerateRefreshToken(user.Username)
	if err != nil {
		c.Error(err)
		return
	}

	payloadRefreshToken := &schema.RefreshToken{
		Token:  refreshToken,
		UserID: userId,
	}

	a.Service.AddRefreshToken(c, payloadRefreshToken, accessToken)
}
