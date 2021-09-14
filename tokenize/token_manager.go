package tokenize

import (
	"os"
	"strconv"
	"time"

	"kurnia123/restfull-api-v1/schema"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(username string) (string, error) {
	age, _ := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_AGE"), 10, 64)
	key := os.Getenv("ACCESS_TOKEN_KEY")

	claims := schema.TokenClaims{
		Username: username,
		Foo:      "bar",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + age,
			Issuer:    "test",
		},
	}

	token, err := GenerateToken(&username, &key, &claims)
	return token, err
}

func GenerateRefreshToken(username string) (string, error) {
	age, _ := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_AGE"), 10, 64)
	key := os.Getenv("REFRESH_TOKEN_KEY")

	claims := schema.TokenClaims{
		Username: username,
		Foo:      "bar",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + (age * 30),
			Issuer:    "test",
		},
	}

	token, err := GenerateToken(&username, &key, &claims)
	return token, err
}

func GenerateToken(username *string, key *string, claims *schema.TokenClaims) (string, error) {
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := sign.SignedString([]byte(*key))

	return token, err
}
