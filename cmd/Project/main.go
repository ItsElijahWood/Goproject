package main

import (
	"log"
)

func main() {
    log.Println("Listening on :8080...")
    err := routes.startServer()
    if err != nil {
        log.Fatal(err)
    }
}
