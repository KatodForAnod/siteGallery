package controller

import (
	"KatodForAnod/siteGallery/internal/models"
	"fmt"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	log.SetLevel(log.ErrorLevel)
}

func TestController_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUser := models.User{}

	mockDB := NewMockDatabase(ctrl)
	mockDB.EXPECT().AddUser(testUser).Return(nil)
	testController := Controller{}
	testController.db = mockDB

	assert.Equal(t, nil, testController.CreateUser(testUser),
		"in that case controller must return nil")
}

func TestGetControllerInstance(t *testing.T) {
}

func TestController_GetImages_CheckReturnValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockDatabase(ctrl)
	mockDB.EXPECT().GetImages(int64(1), int64(10)).Return([]models.ImgMetaData{}, nil)
	testController := Controller{}
	testController.db = mockDB

	images, err := testController.GetImages(1, 10)
	assert.Equal(t, nil, err,
		fmt.Sprintf("error was not expected while GetImages: %s", err))
	assert.Equal(t, []models.ImgMetaData{}, images,
		"if db doesnt have images func must return empty array")
}

func TestController_GetImages_CheckArgs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := make([]models.ImgMetaData, 10)
	mockDB := NewMockDatabase(ctrl)
	mockDB.EXPECT().GetImages(int64(0), int64(10)).Return(expected, nil)

	testController := Controller{}
	testController.db = mockDB

	images, err := testController.GetImages(0, 10)
	assert.Equal(t, nil, err,
		fmt.Sprintf("error was not expected while GetImages: %s", err))
	assert.Equal(t, len(expected), len(images),
		"func must return len of array == limit-offset")
}

func TestController_GetUser(t *testing.T) {
}

func TestController_LoadImage(t *testing.T) {
}

func TestController_PrepareImagesPageErr(t *testing.T) {
	_, err := controller.PrepareImagesPage([]models.ImgMetaData{}, -1, "")
	if err == nil {
		t.Error("func must return error with page id == -1")
	}
}

func TestController_PrepareImagesPage(t *testing.T) {
	_, err := controller.PrepareImagesPage([]models.ImgMetaData{}, 1, "")
	assert.Equal(t, nil, err,
		fmt.Sprintf("error was not expected while GetImages: %s", err))
}
