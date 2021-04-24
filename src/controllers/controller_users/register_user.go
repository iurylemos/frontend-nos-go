package controller_users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
		log.Fatal(erro)
	}

	fmt.Println("Usuário", bytes.NewBuffer(user))

	response, erro := http.Post("http://localhost:5000/usuarios", "application/json", bytes.NewBuffer(user))

	if erro != nil {
		log.Fatal(erro)
	}

	defer response.Body.Close()

	fmt.Println(response.Body)
}
