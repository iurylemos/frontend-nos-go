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

func ReadCookies(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados")

	if erro != nil {
		return nil, erro
	}

	// allocating a "map" empty in memory and attributing this for the values variable
	values := make(map[string]string)

	// decoding cookies in attributing in variable values
	if erro = s.Decode("dados", cookie.Value, &values); erro != nil {
		return nil, erro
	}

	return values, nil
}
