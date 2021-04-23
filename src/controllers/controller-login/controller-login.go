package controller_login

import "net/http"

func LoadScreenLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de Login"))
}
