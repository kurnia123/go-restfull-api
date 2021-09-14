package repository

import (
	"kurnia123/restfull-api-v1/exeption"
	"kurnia123/restfull-api-v1/helper"
	"kurnia123/restfull-api-v1/schema"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (u *UserRepository) Save(db *gorm.DB, usr *schema.User) (string, error) {
	id, err := uuid.NewV4()
	helper.PanicError(err)

	payload := usr
	payload.ID = id.String()

	tx := db.Create(payload)
	tx.Commit()
	return id.String(), tx.Error
}

func (u *UserRepository) Update() {
	panic("not implemented") // TODO: Implement
}

func (u *UserRepository) Delete() {
	panic("not implemented") // TODO: Implement
}

func (u *UserRepository) FindById(c *gin.Context, db *gorm.DB, id string) (*schema.User, error) {
	usr := &schema.User{}

	tx := db.First(&usr, "id = ?", id)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (u *UserRepository) FindByUsername(c *gin.Context, db *gorm.DB, username string) (*schema.User, error) {
	usr := &schema.User{}

	tx := db.First(&usr, "username = ?", username)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (u *UserRepository) FindAll(c *gin.Context, db *gorm.DB) (*[]schema.User, error) {
	usr := &[]schema.User{}

	tx := db.Find(&usr)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func VerifyUserCredential(db *gorm.DB, username string, password string) (string, error) {
	usr := &schema.User{}
	tx := db.Where("username = ? AND password = ?", username, password).First(&usr)

	if tx.Error != nil {
		return "", exeption.NewNotFoundError("Username / Password anda salah")
	}

	return usr.ID, nil
}
