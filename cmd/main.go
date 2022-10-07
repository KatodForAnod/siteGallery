package main

import (
	"KatodForAnod/siteGallery/internal/config"
	"KatodForAnod/siteGallery/internal/controller"
	"KatodForAnod/siteGallery/internal/view"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

const dirName = "logs"

func logInit() error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fileLogName := time.Now().Format("2006-01-02") + ".txt"
	f, err := os.OpenFile(dirName+"/"+fileLogName, os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
	return nil
}

func main() {
	logInit()

	conf, _ := config.LoadConfig()
	contrllr, err := controller.GetControllerInstance(conf)
	if err != nil {
		log.Println(err)
		return
	}
	view.StartHttpServer(conf, contrllr)
}
