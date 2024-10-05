package routes

import (
	"net/http"
)

func SetupRoutes() {
    fs := http.FileServer(http.Dir("../../cmd/Project/static"))
    http.Handle("/", fs)
}
