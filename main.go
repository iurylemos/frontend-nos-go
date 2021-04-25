package main

import (
	"fmt"
	"frontend-nos/src/config"
	"frontend-nos/src/router"
	"frontend-nos/src/utils"
	"log"
	"net/http"
)

func main() {
	config.LoadEnviroment()
	utils.LoadTemplates()
	r := router.Router()

	fmt.Println(config.Port)

	fmt.Printf("Running in port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
