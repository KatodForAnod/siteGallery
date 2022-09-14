package main

import (
	"KatodForAnod/siteGallery/internal/config"
	"KatodForAnod/siteGallery/internal/controller"
	"KatodForAnod/siteGallery/internal/view"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	conf, _ := config.LoadConfig()
	contrllr, err := controller.GetControllerInstance(conf)
	if err != nil {
		log.Println(err)
		return
	}
	view.StartHttpServer(conf, contrllr)
}
