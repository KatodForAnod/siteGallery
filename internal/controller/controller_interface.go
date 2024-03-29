package controller

import "KatodForAnod/siteGallery/internal/models"

type UserController interface {
	CreateUser(user models.User) error
	CreateUserRetId(user models.User) (int64, error)
	GetUser(email string) (models.User, error)
}

type ImageController interface {
	GetImages(offset, limit int64) ([]models.ImgMetaData, error)
	LoadImage(data models.ImgMetaData) error
	PrepareImagesPage(imagesArr []models.ImgMetaData,
		id int, urlBase string) (models.ImagesPageBody, error)
	GetImage(id int64) (models.ImgMetaData, error)
}

type CommonController interface {
	UserController
	ImageController
}
