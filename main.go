package main

import (
	"fmt"
	"frontend-nos/src/routes"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running frontend-nos")

	r := routes.Routes()

	log.Fatal(http.ListenAndServe(":3000", r))
}
