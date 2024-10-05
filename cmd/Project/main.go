package main

import (
	"log"
	"net/http"

	"github.com/ItsElijahWood/Goproject/internal/routes"
)

func main() {
    routes.SetupRoutes()

    log.Println("Listening on :8080...")
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal(err)
    }
}
