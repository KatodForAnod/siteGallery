package view

import (
	"KatodForAnod/siteGallery/internal/controller"
	"KatodForAnod/siteGallery/internal/models"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"strconv"
)

type Handlers struct {
	controller controller.Controller
}

func (h *Handlers) GetImagesPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/tmpls/index.html", "internal/tmpls/imgBlock.html")
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusBadRequest, w)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		idStr = "1"
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusInternalServerError, w)
		return
	} else if id < 1 {
		h.ErrorHandling("wrong page", http.StatusInternalServerError, w)
		return
	}

	limit := 14 //TODO magic numbers?
	offset := (id - 1) * limit

	imagesArr, err := h.controller.GetImages(int64(offset), int64(limit))
	if err != nil {
		log.Println(err)
		h.ErrorHandling(err.Error(), http.StatusInternalServerError, w)
		return
	}

	pageBody, err := h.controller.PrepareImagesPage(imagesArr, id, "/mainPg")
	if err != nil {
		log.Println(err)
		h.ErrorHandling(err.Error(), http.StatusInternalServerError, w)
		return
	}

	pageBody.ImagePlus = template.URL(plusImage)
	pageBody.ImageBackground = template.URL(mainPageBackgroundImage)

	err = tmpl.ExecuteTemplate(w, "MainPage", pageBody)
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusBadRequest, w)
		return
	}
}

func (h *Handlers) LoadImagePageGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/tmpls/index.html", "internal/tmpls/downloadFile.html")
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusBadRequest, w)
		return
	}

	err = tmpl.ExecuteTemplate(w, "MainPage", nil)
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusBadRequest, w)
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
		h.ErrorHandling(err.Error(), http.StatusBadRequest, w)
		return
	}

	if fileHeader.Size > MAX_UPLOAD_SIZE {
		h.ErrorHandling("The uploaded file is too big. Please choose an file that's less than 1MB in size",
			http.StatusBadRequest, w)
		return
	}

	buff := make([]byte, MAX_UPLOAD_SIZE)
	_, err = file.Read(buff)
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusInternalServerError, w)
		return
	}

	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		h.ErrorHandling(err.Error(), http.StatusInternalServerError, w)
		return
	}

	dataEncode := base64.StdEncoding.EncodeToString(buff)
	newImage := models.ImgMetaData{
		FileName:   fileHeader.Filename,
		Tags:       []string{},
		Data:       template.URL(fmt.Sprintf("tmpls:%s;base64,%s", contentType, dataEncode)),
		LoadByUser: "",
	}

	_ = h.controller.LoadImage(newImage)
	http.Redirect(w, r, "/mainPg", http.StatusTemporaryRedirect)
}
