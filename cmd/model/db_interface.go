package model

type Database interface {
	DatabaseImg
	DatabaseUser
}

type DatabaseImg interface {
	AddImage(data ImgMetaData) error
	RemoveImage(id int64) error
	GetImage(id int64) (ImgMetaData, error)
	GetImages(offSet, limit int64) ([]ImgMetaData, error)
}

type DatabaseUser interface {
	AddUser(user User) error
	GetUser(email string) (User, error)
}
