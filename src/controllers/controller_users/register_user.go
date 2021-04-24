package controller_users

import (
	"frontend-nos/src/utils"
	"net/http"
)

func LoadScreenRegister(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "register.html", nil)
}
