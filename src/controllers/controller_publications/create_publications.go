package controller_publications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend-nos/src/config"
	"frontend-nos/src/utils"
	"net/http"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entrou aqui?")
	r.ParseForm()

	publication, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, utils.ErrorAPI{Erro: erro.Error()})
		return
	}

	// creating url of api

	url := fmt.Sprintf("%s/publicacoes", config.API_URL)

	response, erro := utils.MakeRequestWithAuthetication(r, http.MethodPost, url, bytes.NewBuffer(publication))

	if erro != nil {
		utils.ResponseJSON(w, http.StatusInternalServerError, utils.ErrorAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		utils.HandleStatusCodeError(w, response)
		return
	}

	utils.ResponseJSON(w, response.StatusCode, nil)
}
