package routes

import (
	"fmt"
	"net/http"
)

func SetupRoutes() {
	fmt.Println("Routes.go is running") 
	fs := http.FileServer(http.Dir("cmd/Project/static")) // fetch static files from static folder
	http.Handle("/", fs) 
}
