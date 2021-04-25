package controller_users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend-nos/src/config"
	"frontend-nos/src/utils"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// nome := r.FormValue("nome")
	// fmt.Println(nome)

	user, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrorAPI{Erro: erro.Error()})
		return
	}

	// fmt.Println("Usuário", bytes.NewBuffer(user))
	url := fmt.Sprintf("%s/usuarios", config.API_URL)

	response, erro := http.Post(url, "application/json", bytes.NewBuffer(user))

	if erro != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, utils.ErrorAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	// fmt.Println(response.Body)

	if response.StatusCode >= 400 {
		utils.HandleStatusCodeError(w, response)
		return
	}

	utils.ResponseJSON(w, response.StatusCode, nil)
}
