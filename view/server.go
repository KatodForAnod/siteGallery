package view

import (
	"log"
	"net/http"
	"siteGallery/config"
	"siteGallery/controller"
)

func StartHttpServer(loadedConf config.Config, controller controller.Controller) error {
	//addr := loadedConf.SvConfig.Host + ":" + loadedConf.SvConfig.Port
	//fmt.Println("Server is listening...", addr)

	handlers := Handlers{controller: controller}
	mux := http.NewServeMux()

	mux.HandleFunc("/mainPg", handlers.GetImagesPage)
	mux.HandleFunc("/loadImg", handlers.LoadImagePageGet)
	mux.HandleFunc("/loadImage2", handlers.LoadImagePagePost)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	return nil
}
