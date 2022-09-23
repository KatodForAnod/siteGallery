package view

import (
	"KatodForAnod/siteGallery/internal/config"
	"KatodForAnod/siteGallery/internal/controller"
	_ "embed"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//go:embed  images/loginRegistrationImage
var loginRegistrationImage string

//go:embed  images/plusImage
var plusImage string

//go:embed  images/mainPageBackgroundImage
var mainPageBackgroundImage string

func StartHttpServer(loadedConf config.Config, controller controller.Controller) error {
	addr := loadedConf.SvConfig.Host + ":" + loadedConf.SvConfig.Port
	fmt.Println("Server is listening... ", addr)

	handlers := Handlers{controller: controller}
	mux := http.NewServeMux()

	mux.HandleFunc("/mainPg", handlers.GetImagesPage)
	mux.HandleFunc("/loadImg", handlers.CheckAuth(handlers.LoadImagePageGet))
	mux.HandleFunc("/postImage", handlers.CheckAuth(handlers.LoadImagePagePost))
	mux.HandleFunc("/registration", handlers.Registration)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/logout", handlers.Logout)

	mux.HandleFunc("/log", handlers.GetLoginPage)
	mux.HandleFunc("/reg", handlers.GetRegistrationPage)

	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
	return nil
}
