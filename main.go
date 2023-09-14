package main

import (
	"fmt"
	"go/postgres-go/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 5000")

	err := http.ListenAndServe(":5000", r)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
