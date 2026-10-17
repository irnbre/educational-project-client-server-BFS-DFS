package main

import (
	"net/http"
	"petproj/server/handler"
)

func NewRouter() *http.ServeMux {

	rout := http.NewServeMux()

	rout.HandleFunc("POST /search", handler.AlgoHandler)

	return rout
}
