package main

import (
	"log"
	"net/http"

	iform "github.com/ItsElijahWood/Goproject/cmd/Project/data"
	database "github.com/ItsElijahWood/Goproject/internal/database"
	routes "github.com/ItsElijahWood/Goproject/internal/routes"
)

func main() {
    // imported thru internals
    database.ConnectToMongoDB() 
	routes.SetupRoutes() 

    // handle forms db
    http.HandleFunc("/submit", iform.IForm)

	log.Println("Listening on :80...")
	err := http.ListenAndServe(":80", nil) 
	if err != nil {
		log.Fatal(err)
	}
}
