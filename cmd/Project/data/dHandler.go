package dhandler

import (
	"net/http"

	iform "github.com/ItsElijahWood/Goproject/cmd/Project/data/handler"
)

func DHandler() {
 	http.HandleFunc("/submit", iform.IForm)
}