package controller_pages

import (
	"fmt"
	"frontend-nos/src/config"
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

	fmt.Println("Entrou aqui?", response.StatusCode, erro)

	utils.ExecTemplate(w, "home.html", nil)
}
