package view

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"siteGallery/cmd/controller"
	"siteGallery/cmd/model"
)

type Handlers struct {
	controller controller.Controller
}

func (h *Handlers) GetImagesPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("cmd/data/index.tmpl", "cmd/data/imgBlock.tmpl")
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

func (h *Handlers) LoadImagePageGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("cmd/data/index.tmpl", "cmd/data/downloadFile.tmpl")
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
	newImage := model.ImgMetaData{
		FileName:   fileHeader.Filename,
		Tags:       []string{},
		Data:       template.URL(fmt.Sprintf("data:%s;base64,%s", contentType, dataEncode)),
		LoadByUser: "",
	}

	_ = h.controller.LoadImage(newImage)
	http.Redirect(w, r, "/mainPg", http.StatusTemporaryRedirect)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	user := model.User{
		User:     username,
		Email:    email,
		PassHash: string(hashedPass),
	}

	if err := h.controller.CreateUser(user); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	token, err := CreateToken(email)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", token)
}

func (h *Handlers) Auth(w http.ResponseWriter, r *http.Request) {
	//username :=	r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	user, err := h.controller.GetUser(email)
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	if user.Email != email {
		log.Println(err)
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	} else if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash),
		[]byte(password)); err != nil {
		log.Println(err)
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := CreateToken(email)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", token)
}

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Del("Authorization")
	http.Redirect(w, r, "/mainPg", http.StatusTemporaryRedirect)
}
