package main

import (
	"log"

	"github.com/ItsElijahWood/Goproject/internal/routes"
)

func main() {
    log.Println("Listening on :8080...")
    err := routes.StartServer()
    if err != nil {
        log.Fatal(err)
    }
}
