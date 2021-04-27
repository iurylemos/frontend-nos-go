package controller_pages

import (
	"encoding/json"
	"fmt"
	"frontend-nos/src/config"
	"frontend-nos/src/models"
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
	url := fmt.Sprintf("%s/publicacoes", config.API_URL)

	response, erro := utils.MakeRequestWithAuthetication(r, http.MethodGet, url, nil)

	if erro != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, utils.ErrorAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		utils.HandleStatusCodeError(w, response)
		return
	}

	var publications []models.Publication

	if erro = json.NewDecoder(response.Body).Decode(&publications); erro != nil {
		utils.ResponseJSON(w, http.StatusUnprocessableEntity, utils.ErrorAPI{Erro: erro.Error()})
		return
	}

	utils.ExecTemplate(w, "home.html", publications)
}
