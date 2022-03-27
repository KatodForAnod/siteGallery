package controller

import (
	"KatodForAnod/siteGallery/cmd/config"
	model2 "KatodForAnod/siteGallery/cmd/model"
	"KatodForAnod/siteGallery/cmd/model/db"
	"errors"
	"html/template"
	"log"
	"strconv"
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

	arr, err := c.db.GetImages(offset, limit)
	if err != nil {
		log.Println(err)
		return []model2.ImgMetaData{}, err
	}

	if len(arr) == 0 {
		return []model2.ImgMetaData{}, errors.New("no images")
	}

	for len(arr) < int(limit) {
		arr = append(arr, model2.ImgMetaData{})
	}

	return arr, nil
}

func (c *Controller) PrepareImagesPage(imagesArr []model2.ImgMetaData,
	id int, urlBase string) (model2.ImagesPageBody, error) {
	if id < 1 {
		return model2.ImagesPageBody{}, errors.New("wrong input data")
	}

	var outputData model2.ImagesPageBody
	outputData.ImagesArr = imagesArr

	var nextId, prevId string

	nextId = strconv.Itoa(id + 1)
	if id != 1 {
		prevId = strconv.Itoa(id - 1)
	} else {
		prevId = strconv.Itoa(id)
	}

	outputData.PageNext = template.URL(urlBase + "?id=" + nextId)
	outputData.PagePrev = template.URL(urlBase + "?id=" + prevId)

	return outputData, nil
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
