package cookies

import (
	"frontend-nos/src/config"
	"net/http"

	"github.com/gorilla/securecookie"
)

// package used to codify and de decode the information which can stay save in the browser

var s *securecookie.SecureCookie

// function configure it is used to configure the environment variables make creation of secure cookie
func ConfigureCookie() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func SaveCookie(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"ID":    ID,
		"token": token,
	}

	dataCodify, erro := s.Encode("dados", data)

	if erro != nil {
		return erro
	}

	// puting cookie in the browser
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dataCodify,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
