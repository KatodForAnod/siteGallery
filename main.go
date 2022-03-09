package main

import (
	"log"
	"siteGallery/config"
	"siteGallery/controller"
	"siteGallery/view"
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
