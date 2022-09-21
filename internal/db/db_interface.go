package db

import "KatodForAnod/siteGallery/internal/models"

type Database interface {
	DatabaseImg
	DatabaseUser
}

type DatabaseImg interface {
	AddImage(data models.ImgMetaData) error
	RemoveImage(id int64) error
	GetImage(id int64) (models.ImgMetaData, error)
	GetImages(offSet, limit int64) ([]models.ImgMetaData, error)
}

type DatabaseUser interface {
	AddUser(user models.User) error
	GetUser(email string) (models.User, error)
}
