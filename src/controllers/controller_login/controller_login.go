package controller_login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend-nos/src/utils"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrorAPI{Erro: erro.Error()})
		return
	}

	response, erro := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(user))

	if erro != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrorAPI{Erro: erro.Error()})
		return
	}

	token, _ := ioutil.ReadAll(response.Body)

	fmt.Println(response.StatusCode, string(token))
}
