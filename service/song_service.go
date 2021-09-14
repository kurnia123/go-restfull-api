package service

import (
	"kurnia123/restfull-api-v1/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SongService struct {
	SongRepository repository.SongRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewSongService(songRepository repository.SongRepository, DB *gorm.DB, validation *validator.Validate) SongService {
	return SongService{
		SongRepository: songRepository,
		DB:             DB,
		Validate:       validation,
	}
}

func (song *SongService) Save() {
	panic("not implemented") // TODO: Implement
}

func (song *SongService) Update() {
	panic("not implemented") // TODO: Implement
}

func (song *SongService) Delete() {
	panic("not implemented") // TODO: Implement
}

func (song *SongService) FindById() {
	panic("not implemented") // TODO: Implement
}

func (song *SongService) FindAll() {
	panic("not implemented") // TODO: Implement
}
