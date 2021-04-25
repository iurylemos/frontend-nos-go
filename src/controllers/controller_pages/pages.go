package controller_pages

import (
	"frontend-nos/src/utils"
	"net/http"
)

func LoadScreenLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}

func LoadScreenRegister(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "register.html", nil)
}

func LoadScreenMain(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "home.html", nil)
}
