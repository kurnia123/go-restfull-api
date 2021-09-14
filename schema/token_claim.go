package schema

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	Username string `json:"username"`
	Foo      string `json:"foo"`
	jwt.StandardClaims
}
