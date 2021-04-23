package controller_login

import (
	"frontend-nos/src/utils"
	"net/http"
)

func LoadScreenLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}
