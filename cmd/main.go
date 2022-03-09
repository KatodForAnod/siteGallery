package main

import (
	"log"
	"siteGallery/cmd/config"
	"siteGallery/cmd/controller"
	"siteGallery/cmd/view"
)

func main() {
	conf, _ := config.LoadConfig()
	contrllr, err := controller.GetControllerInstance(conf)
	if err != nil {
		log.Println(err)
		return
	}
	view.StartHttpServer(conf, contrllr)
}
