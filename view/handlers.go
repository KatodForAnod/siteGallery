package view

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"siteGallery/controller"
)

type Handlers struct {
	controller controller.Controller
}

func (h Handlers) GetImagesPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("data/index.tmpl", "data/imgBlock.tmpl")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	imagesArr, err := h.controller.GetImages(0, 5)
	if err != nil {
		log.Println(err)
		http.Error(w, "server err", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "MainPage", imagesArr)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (h Handlers) LoadImagePageGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("data/index.tmpl", "data/downloadFile.tmpl")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(w, "MainPage", nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (h Handlers) LoadImagePagePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	model, header, _ := r.FormFile("photo")
	fmt.Println(model, header)
}
