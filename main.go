package main

import (
	"golang-rest-api/config"
	"golang-rest-api/router"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	r := router.Router()
	log.Println("Server started on port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}
