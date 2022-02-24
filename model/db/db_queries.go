package db

import "siteGallery/model"

func (p postgreSQl) AddImage(data model.ImgMetaData) error {
	panic("implement me")
}

func (p postgreSQl) RemoveImage(id int64) error {
	panic("implement me")
}

func (p postgreSQl) GetImage(id int64) (model.ImgMetaData, error) {
	panic("implement me")
}

func (p postgreSQl) GetImages(offSet, limit int64) ([]model.ImgMetaData, error) {
	panic("implement me")
}

func (p postgreSQl) AddUser(user model.User) error {
	panic("implement me")
}

func (p postgreSQl) GetUser(email string) (model.User, error) {
	panic("implement me")
}
