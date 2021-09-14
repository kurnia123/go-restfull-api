package service

import (
	"kurnia123/restfull-api-v1/exeption"
	"kurnia123/restfull-api-v1/repository"
	"kurnia123/restfull-api-v1/schema"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserService struct {
	Repository repository.UserRepository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validation *validator.Validate) UserService {
	return UserService{
		Repository: userRepository,
		DB:         DB,
		Validate:   validation,
	}
}

func (u *UserService) Create(c *gin.Context, usr *schema.User) {
	id, err := u.Repository.Save(u.DB, usr)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"satus":   "posted",
		"message": "success membuat user",
		"id":      id,
	})
}

func (u *UserService) FindByUsername(c *gin.Context, username string) {
	usr, err := u.Repository.FindByUsername(c, u.DB, username)
	if err != nil {
		c.Error(exeption.NewNotFoundError(err.Error()))
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status":  "Finded",
		"message": "seuccess find user",
		"data":    usr,
	})
}

func (u *UserService) FindById(c *gin.Context, id string) {
	usr, err := u.Repository.FindById(c, u.DB, id)
	if err != nil {
		c.Error(exeption.NewNotFoundError(err.Error()))
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status":  "Finded",
		"message": "seuccess find user",
		"data":    usr,
	})
}

func (u *UserService) FindAll(c *gin.Context) {
	usr, err := u.Repository.FindAll(c, u.DB)
	if err != nil {
		c.Error(exeption.NewNotFoundError(err.Error()))
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status":  "Finded",
		"message": "seuccess find user",
		"data":    usr,
	})
}

func VerifyUser(db *gorm.DB, username string, password string) (string, error) {
	id, err := repository.VerifyUserCredential(db, username, password)
	if err != nil {
		return id, err
	}

	return id, nil
}
