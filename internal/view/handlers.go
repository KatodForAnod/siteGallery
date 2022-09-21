package view

import (
	"KatodForAnod/siteGallery/internal/controller"
	"KatodForAnod/siteGallery/internal/models"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Handlers struct {
	controller controller.Controller
}

func (h *Handlers) GetImagesPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("cmd/data/index.html", "cmd/data/imgBlock.tmpl")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		idStr = "1"
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "server err", http.StatusInternalServerError)
		return
	}

	limit := 14
	offset := (id - 1) * limit

	imagesArr, err := h.controller.GetImages(int64(offset), int64(limit))
	if err != nil {
		log.Println(err)
		http.Error(w, "server err", http.StatusInternalServerError)
		return
	}

	pageBody, err := h.controller.PrepareImagesPage(imagesArr, id, "/mainPg")
	if err != nil {
		log.Println(err)
		http.Error(w, "server err", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "MainPage", pageBody)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (h *Handlers) LoadImagePageGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("cmd/data/index.html", "cmd/data/downloadFile.tmpl")
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

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB
func (h *Handlers) LoadImagePagePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if fileHeader.Size > MAX_UPLOAD_SIZE {
		http.Error(w,
			"The uploaded file is too big. Please choose an file that's less than 1MB in size",
			http.StatusInternalServerError)
		return
	}

	buff := make([]byte, MAX_UPLOAD_SIZE)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	dataEncode := base64.StdEncoding.EncodeToString(buff)
	newImage := models.ImgMetaData{
		FileName:   fileHeader.Filename,
		Tags:       []string{},
		Data:       template.URL(fmt.Sprintf("data:%s;base64,%s", contentType, dataEncode)),
		LoadByUser: "",
	}

	_ = h.controller.LoadImage(newImage)
	http.Redirect(w, r, "/mainPg", http.StatusTemporaryRedirect)
}
