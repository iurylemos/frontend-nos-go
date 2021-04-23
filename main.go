package main

import (
	"fmt"
	"frontend-nos/src/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running frontend-nos")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":3000", r))
}
