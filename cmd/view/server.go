package view

import (
	"log"
	"net/http"
	"siteGallery/cmd/config"
	"siteGallery/cmd/controller"
)

func StartHttpServer(loadedConf config.Config, controller controller.Controller) error {
	//addr := loadedConf.SvConfig.Host + ":" + loadedConf.SvConfig.Port
	//fmt.Println("Server is listening...", addr)

	handlers := Handlers{controller: controller}
	mux := http.NewServeMux()

	mux.HandleFunc("/mainPg", handlers.GetImagesPage)
	mux.HandleFunc("/loadImg", handlers.LoadImagePageGet)
	mux.HandleFunc("/loadImage2", handlers.LoadImagePagePost)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/auth", handlers.Auth)
	mux.HandleFunc("/logout", handlers.Logout)

	mux.HandleFunc("/log", handlers.GetLoginPage)
	mux.HandleFunc("/reg", handlers.GetRegistrationPage)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	return nil
}
