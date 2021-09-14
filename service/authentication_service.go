package service

import (
	"kurnia123/restfull-api-v1/repository"
	"kurnia123/restfull-api-v1/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthService struct {
	Repository repository.AuthRepository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewAuthService(
	auth repository.AuthRepository,
	db *gorm.DB,
	validate *validator.Validate) AuthService {

	return AuthService{
		Repository: auth,
		DB:         db,
		Validate:   validate,
	}
}

func (auth *AuthService) AddRefreshToken(c *gin.Context, payload *schema.RefreshToken, accessToken string) {
	err := auth.Repository.Save(c, auth.DB, payload)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": payload.Token,
	})
}

func (auth *AuthService) Update() {
	panic("not implemented") // TODO: Implement
}

func (auth *AuthService) Delete() {
	panic("not implemented") // TODO: Implement
}

func (auth *AuthService) FindById() {
	panic("not implemented") // TODO: Implement
}

func (auth *AuthService) FindAll() {
	panic("not implemented") // TODO: Implement
}

func (auth *AuthService) AuthenticationUser(c *gin.Context, username string, password string) string {
	res, err := VerifyUser(auth.DB, username, password)

	if err != nil {
		c.Error(err)
		return res
	}

	return res
}
