package view

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"siteGallery/cmd/model"
)

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
