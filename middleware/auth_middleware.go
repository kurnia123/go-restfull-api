package middleware

import (
	"fmt"
	"os"
	"strings"

	"kurnia123/restfull-api-v1/exeption"
	"kurnia123/restfull-api-v1/schema"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	if reqToken == "" {
		c.Error(exeption.NewUnAuthorizedError("Not Found Token Bearer"))
		c.Abort()
		return
	}

	tokenString := strings.Split(reqToken, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString[1], &schema.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
	})

	// if claims, ok := token.Claims.(*schema.TokenClaims); ok && token.Valid {
	// 	fmt.Println("claims :", claims)
	// }

	if token == nil || err != nil {
		c.Error(exeption.NewUnAuthorizedError("Not Authorized"))
		c.Abort()
		return
	}
}
