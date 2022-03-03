package view

import (
	"fmt"
	"log"
	"net/http"
	"siteGallery/config"
	"siteGallery/controller"
)

func StartHttpServer(loadedConf config.Config, controller controller.Controller) error {
	addr := loadedConf.SvConfig.Host + ":" + loadedConf.SvConfig.Port
	fmt.Println("Server is listening...", addr)

	handlers := Handlers{controller: controller}
	mux := http.NewServeMux()

	mux.HandleFunc("/test", handlers.GetImagesPage)
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
	return nil
}
