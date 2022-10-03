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
	controller controller.CommonController
}

func (h *Handlers) GetImagesPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"internal/tmpls/index.html",
		"internal/tmpls/imgBlock.html",
	}
	if h.CheckAuth(r) {
		files = append(files, "internal/tmpls/indexHeaderUnLogin.html")
	} else {
		files = append(files, "internal/tmpls/indexHeaderLogin.html")
	}

	tmpl, err := template.ParseFiles(files...)
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
	// TODO fix background image disappear
	files := []string{
		"internal/tmpls/index.html",
		"internal/tmpls/downloadFile.html",
	}
	if h.CheckAuth(r) {
		files = append(files, "internal/tmpls/indexHeaderUnLogin.html")
	} else {
		files = append(files, "internal/tmpls/indexHeaderLogin.html")
	}

	tmpl, err := template.ParseFiles(files...)
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

	userId, err := h.GetUserId(r)
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusInternalServerError, w)
		return
	}

	tag1 := r.FormValue("tag1")
	tag2 := r.FormValue("tag2")
	tag3 := r.FormValue("tag3")
	dataEncode := base64.StdEncoding.EncodeToString(buff)
	newImage := models.ImgMetaData{
		FileName: fileHeader.Filename,
		Tags:     []string{tag1, tag2, tag3},
		Data:     template.URL(fmt.Sprintf("data:%s;base64,%s", contentType, dataEncode)),
		UserId:   userId,
	}

	err = h.controller.LoadImage(newImage)
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusInternalServerError, w)
		return
	}

	http.Redirect(w, r, "/mainPg", http.StatusTemporaryRedirect)
}
