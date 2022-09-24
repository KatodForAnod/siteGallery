package controller

import (
	"KatodForAnod/siteGallery/internal/config"
	"KatodForAnod/siteGallery/internal/db"
	"KatodForAnod/siteGallery/internal/db/postgres"
	"KatodForAnod/siteGallery/internal/models"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"strconv"
	"sync"
)

var (
	controller Controller
	sc         sync.Once
)

type Controller struct {
	db db.Database
}

func GetControllerInstance(config config.Config) (CommonController, error) {
	var err error
	sc.Do(func() {
		var dbModel db.Database
		dbModel, err = postgres.GetPostgresSQlConn(config)
		if err != nil {
			log.Errorln(err)
			return
		}
		controller.db = dbModel
	})

	return &controller, err
}

func (c *Controller) GetImages(offset, limit int64) ([]models.ImgMetaData, error) {
	log.Println("GetImages controller")

	arr, err := c.db.GetImages(offset, limit)
	if err != nil {
		return []models.ImgMetaData{}, fmt.Errorf("GetImages err: %s", err)
	}

	if len(arr) == 0 {
		return []models.ImgMetaData{}, nil
	}

	for len(arr) < int(limit) {
		arr = append(arr, models.ImgMetaData{})
	}

	return arr, nil
}

func (c *Controller) PrepareImagesPage(imagesArr []models.ImgMetaData,
	id int, urlBase string) (models.ImagesPageBody, error) {
	if id < 1 {
		return models.ImagesPageBody{}, errors.New("PrepareImagesPage: wrong input data")
	}

	var outputData models.ImagesPageBody
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

func (c *Controller) LoadImage(data models.ImgMetaData) error {
	//log.Println("LoadImage controller")
	return c.db.AddImage(data)
}

func (c *Controller) CreateUser(user models.User) error {
	//log.Println("Create user")
	return c.db.AddUser(user)
}

func (c *Controller) GetUser(email string) (models.User, error) {
	//log.Println("Get user")
	return c.db.GetUser(email)
}
