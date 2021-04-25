package middlewares

import (
	"frontend-nos/src/cookies"
	"log"
	"net/http"
)

//Logger written the information of request in prompt command
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// verifying if the cookies being saved in cookies of browser
func Autheticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// cookies.ReadCookies
		// if this function return error let's redirect the user to screen the login

		if _, erro := cookies.ReadCookies(r); erro != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next(w, r)
	}
}
