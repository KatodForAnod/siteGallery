package view

import (
	"KatodForAnod/siteGallery/internal/models"
	_ "embed"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

//go:embed  images/loginRegistrationImage
var loginRegistrationImage string

type LoginRegistrationPage struct {
	ActionLink string
	ActionName string

	ImageBackground template.URL
}

func (h *Handlers) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/tmpls/auth.html")
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusBadRequest, w)
		return
	}

	f := LoginRegistrationPage{ActionLink: "/login", ActionName: "Войти",
		ImageBackground: template.URL(loginRegistrationImage)}

	err = tmpl.Execute(w, f)
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusBadRequest, w)
		return
	}
}

func (h *Handlers) GetRegistrationPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/tmpls/auth.html")
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusBadRequest, w)
		return
	}

	f := LoginRegistrationPage{ActionLink: "/registration",
		ActionName: "Зарегистрироваться", ImageBackground: template.URL(loginRegistrationImage)}

	err = tmpl.Execute(w, f)
	if err != nil {
		h.ErrorHandling(err.Error(), http.StatusBadRequest, w)
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
		h.ErrorHandling(err.Error(), http.StatusInternalServerError, w)
		return
	}

	user := models.User{
		User:     username,
		Email:    email,
		PassHash: string(hashedPass),
	}

	if err := h.controller.CreateUser(user); err != nil {
		log.Println(err)
		h.ErrorHandling(err.Error(), http.StatusInternalServerError, w)
		return
	}

	token, err := CreateToken(email)
	if err != nil {
		log.Println(err)
		h.ErrorHandling(err.Error(), http.StatusInternalServerError, w)
		return
	}

	authCookie := http.Cookie{Name: "x-token", Value: token}
	http.SetCookie(w, &authCookie)
	//w.Header().Set("Authorization", token)
	http.Redirect(w, r, "/mainPg", http.StatusTemporaryRedirect)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	//username :=	r.FormValue("username")
	password := r.FormValue("auth_pass")
	email := r.FormValue("auth_email")

	user, err := h.controller.GetUser(email)
	if err != nil {
		log.Println(err)
		h.ErrorHandling("invalid email or password", http.StatusUnauthorized, w)
		return
	}

	if user.Email != email {
		log.Println(err)
		h.ErrorHandling("invalid email or password", http.StatusUnauthorized, w)
		return
	} else if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash),
		[]byte(password)); err != nil {
		log.Println(err)
		h.ErrorHandling("invalid email or password", http.StatusUnauthorized, w)
		return
	}

	token, err := CreateToken(email)
	if err != nil {
		log.Println(err)
		h.ErrorHandling(err.Error(), http.StatusInternalServerError, w)
		return
	}

	authCookie := http.Cookie{Name: "x-token", Value: token} // x-token to constant
	http.SetCookie(w, &authCookie)
	//w.Header().Set("Authorization", token)
	http.Redirect(w, r, "/mainPg", http.StatusTemporaryRedirect)
}

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Del("Authorization")
	http.SetCookie(w, &http.Cookie{Name: "x-token"})
	http.Redirect(w, r, "/mainPg", http.StatusTemporaryRedirect)
}

func (h *Handlers) CheckAuth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, cookie := range r.Cookies() {
			if cookie.Name == "x-token" {
				verifyToken, err := VerifyToken(cookie.Value)
				if err != nil {
					h.ErrorHandling("pls re-login", http.StatusUnauthorized, w)
					return
				}
				if !verifyToken.Valid {
					h.ErrorHandling("pls re-login", http.StatusUnauthorized, w)
					return
				}

				f(w, r)
				return
			}
		}

		h.ErrorHandling("pls login", http.StatusUnauthorized, w)
		return
	}
}

func (h *Handlers) ErrorHandling(errorMsg string, statusCode int, w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("internal/tmpls/errorHandling.html")
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	err = tmpl.Execute(w, errorMsg)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
}
