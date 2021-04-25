package main

import (
	"fmt"
	"frontend-nos/src/config"
	"frontend-nos/src/cookies"
	"frontend-nos/src/router"
	"frontend-nos/src/utils"
	"log"
	"net/http"
)

// func init() {
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))

// 	fmt.Println(hashKey)

// 	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))

// 	fmt.Println(blockKey)
// }

func main() {
	config.LoadEnviroment()
	cookies.ConfigureCookie()
	utils.LoadTemplates()
	r := router.Router()

	fmt.Printf("Running in port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
