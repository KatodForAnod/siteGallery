package controller

import (
	"KatodForAnod/siteGallery/internal/models"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testController = Controller{db: database}
var database fakeDataBase

func init() {
	log.SetLevel(log.ErrorLevel)
}

type fakeDataBase struct{}

func (d fakeDataBase) AddImage(data models.ImgMetaData) error {
	return nil
}

func (d fakeDataBase) RemoveImage(id int64) error {
	return nil
}

func (d fakeDataBase) GetImage(id int64) (models.ImgMetaData, error) {
	return models.ImgMetaData{}, nil
}

func (d fakeDataBase) GetImages(offSet, limit int64) ([]models.ImgMetaData, error) {
	if offSet >= 1 {
		return []models.ImgMetaData{}, nil
	}

	return []models.ImgMetaData{models.ImgMetaData{}}, nil
}

func (d fakeDataBase) AddUser(user models.User) error {
	return nil
}

func (d fakeDataBase) GetUser(email string) (models.User, error) {
	return models.User{}, nil
}

func TestController_CreateUser(t *testing.T) {
	assert.Equal(t, nil, testController.CreateUser(models.User{}))
}

func TestGetControllerInstance(t *testing.T) {
}

func TestController_GetImages(t *testing.T) {
	images, _ := testController.GetImages(1, 10)
	assert.Equal(t, []models.ImgMetaData{}, images,
		"if db doesnt have images func must return empty array")
}

func TestController_GetImages2(t *testing.T) {
	images, _ := testController.GetImages(0, 10)
	expected := make([]models.ImgMetaData, 10)
	assert.Equal(t, len(expected), len(images),
		"func must return len of array == limit-offset")
}

func TestController_GetUser(t *testing.T) {
}

func TestController_LoadImage(t *testing.T) {
}

func TestController_PrepareImagesPage(t *testing.T) {
	_, err := controller.PrepareImagesPage([]models.ImgMetaData{}, -1, "")

	if err == nil {
		t.Error("func must return error with page id == -1")
	}
}

func TestController_PrepareImagesPage2(t *testing.T) {
	_, err := controller.PrepareImagesPage([]models.ImgMetaData{}, 1, "")
	assert.Equal(t, nil, err)
}
