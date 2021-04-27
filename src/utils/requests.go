package utils

import (
	"frontend-nos/src/cookies"
	"io"
	"net/http"
)

// it is used to place token in request
func MakeRequestWithAuthetication(r *http.Request, method, url string, dados io.Reader) (*http.Response, error) {

	// in first parameter the function i am received a request
	// but i am doing this to received the cookies of request

	// i am created the request but i am not doing the request to api
	request, erro := http.NewRequest(method, url, dados)

	if erro != nil {
		return nil, erro
	}

	cookie, _ := cookies.ReadCookies(r)

	request.Header.Add("Authorization", "Bearer "+cookie["token"]) // Bearer Token

	// creating client to do a request of facts

	client := &http.Client{}

	response, erro := client.Do(request)

	if erro != nil {
		return nil, erro
	}

	return response, nil
}
