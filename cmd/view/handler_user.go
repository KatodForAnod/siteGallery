package view

import (
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"siteGallery/cmd/model"
)

type LoginRegistrationPage struct {
	ActionLink string
	ActionName string
}

func (h *Handlers) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("cmd/data/auth.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	f := LoginRegistrationPage{ActionLink: "/login", ActionName: "Войти"}

	err = tmpl.Execute(w, f)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (h *Handlers) GetRegistrationPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("cmd/data/auth.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	f := LoginRegistrationPage{ActionLink: "/registration", ActionName: "Зарегистрироваться"}

	err = tmpl.Execute(w, f)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (h *Handlers) Registration(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("auth_pass")
	email := r.FormValue("auth_email")

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

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	//username :=	r.FormValue("username")
	password := r.FormValue("auth_pass")
	email := r.FormValue("auth_email")

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
