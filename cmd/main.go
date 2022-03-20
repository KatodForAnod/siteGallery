package main

import (
	"KatodForAnod/siteGallery/cmd/config"
	"KatodForAnod/siteGallery/cmd/controller"
	"KatodForAnod/siteGallery/cmd/view"
	"log"
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
