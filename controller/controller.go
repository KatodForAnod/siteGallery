package controller

import (
	"log"
	"siteGallery/config"
	"siteGallery/model"
	"siteGallery/model/db"
	"sync"
)

var (
	controller Controller
	sc         sync.Once
)

type Controller struct {
	db model.Database
}

func GetControllerInstance(config config.Config) (Controller, error) {
	var err error
	sc.Do(func() {
		var dbModel model.Database
		dbModel, err = db.GetPostgresSQlConn(config)
		if err != nil {
			log.Println(err)
			return
		}
		controller.db = dbModel
	})

	return controller, err
}

func (c *Controller) GetImages(offset, limit int64) ([]model.ImgMetaData, error) {
	log.Println("GetImages controller")
	return c.db.GetImages(offset, limit)
}

func (c *Controller) LoadImage(data model.ImgMetaData) error {
	log.Println("LoadImage controller")
	return c.db.AddImage(data)
}
