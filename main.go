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

func main() {
	config.LoadEnviroment()
	cookies.ConfigureCookie()
	utils.LoadTemplates()
	r := router.Router()

	fmt.Printf("Running in port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
