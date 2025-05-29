package main

import (
	"emotion-detector/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.HomeHandler)

	log.Println("Servidor iniciado em http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
