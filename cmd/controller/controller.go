package controller

import (
	"KatodForAnod/siteGallery/cmd/config"
	model2 "KatodForAnod/siteGallery/cmd/model"
	"KatodForAnod/siteGallery/cmd/model/db"
	"log"
	"sync"
)

var (
	controller Controller
	sc         sync.Once
)

type Controller struct {
	db model2.Database
}

func GetControllerInstance(config config.Config) (Controller, error) {
	var err error
	sc.Do(func() {
		var dbModel model2.Database
		dbModel, err = db.GetPostgresSQlConn(config)
		if err != nil {
			log.Println(err)
			return
		}
		controller.db = dbModel
	})

	return controller, err
}

func (c *Controller) GetImages(offset, limit int64) ([]model2.ImgMetaData, error) {
	log.Println("GetImages controller")
	return c.db.GetImages(offset, limit)
}

func (c *Controller) LoadImage(data model2.ImgMetaData) error {
	log.Println("LoadImage controller")
	return c.db.AddImage(data)
}

func (c *Controller) CreateUser(user model2.User) error {
	log.Println("Create user")
	return c.db.AddUser(user)
}

func (c *Controller) GetUser(email string) (model2.User, error) {
	log.Println("Get user")
	return c.db.GetUser(email)
}
