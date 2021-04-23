package main

import (
	"fmt"
	"frontend-nos/src/router"
	"frontend-nos/src/utils"
	"log"
	"net/http"
)

func main() {
	utils.LoadTemplates()
	r := router.Router()

	fmt.Println("Running in port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
