package routes

import (
	"net/http"
)

func startServer() error {
    fs := http.FileServer(http.Dir("../../cmd/Project/static"))
    http.Handle("/", fs)
    return http.ListenAndServe(":80", nil)
}
