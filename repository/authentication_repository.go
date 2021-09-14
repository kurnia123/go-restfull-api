package repository

import (
	"kurnia123/restfull-api-v1/helper"
	"kurnia123/restfull-api-v1/schema"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
}

func NewAuthRepository() AuthRepository {
	return AuthRepository{}
}

func (auth *AuthRepository) Save(c *gin.Context, db *gorm.DB, payload *schema.RefreshToken) error {
	id, err := uuid.NewV4()
	helper.PanicError(err)

	data := payload
	data.ID = id.String()

	tx := db.Create(data)
	err = tx.Error
	if err != nil {
		return err
	}

	return nil
}

func (auth *AuthRepository) Update() {
	panic("not implemented") // TODO: Implement
}

func (auth *AuthRepository) Delete() {
	panic("not implemented") // TODO: Implement
}

func (auth *AuthRepository) FindById() {
	panic("not implemented") // TODO: Implement
}

func (auth *AuthRepository) FindAll() {
	panic("not implemented") // TODO: Implement
}
