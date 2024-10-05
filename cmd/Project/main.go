package main

import (
	"log"
	"net/http"

	"github.com/ItsElijahWood/Goproject/internal/routes"
)

func main() {
	log.Println("Setting up routes...")

	routes.SetupRoutes() // importing routes from internals

	log.Println("Listening on :80...")
	err := http.ListenAndServe(":80", nil) 
	if err != nil {
		log.Fatal(err)
	}
}
