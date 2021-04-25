package controller_login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend-nos/src/config"
	"frontend-nos/src/models"
	"frontend-nos/src/utils"
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

	url := fmt.Sprintf("%s/login", config.API_URL)

	response, erro := http.Post(url, "application/json", bytes.NewBuffer(user))

	if erro != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrorAPI{Erro: erro.Error()})
		return
	}

	// token, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(response.StatusCode, string(token))

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		utils.HandleStatusCodeError(w, response)
		return
	}

	var dataAuth models.DataAuth

	if erro = json.NewDecoder(response.Body).Decode(&dataAuth); erro != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrorAPI{Erro: erro.Error()})
		return
	}

	// salving data user in cookies

	utils.ResponseJSON(w, http.StatusOK, nil)

}
