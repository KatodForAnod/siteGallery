package controller

import (
	"log"
	"siteGallery/model"
)

type Controller struct {
	db model.Database
}

func (c *Controller) GetImages(offset, limit int64) ([]model.ImgMetaData, error) {
	log.Println("GetImages controller")
	return c.db.GetImages(offset, limit)
}

func (c *Controller) LoadImage(data model.ImgMetaData) error {
	log.Println("LoadImage controller")
	return c.db.AddImage(data)
}
