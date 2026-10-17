package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := NewRouter()

	server := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Сервер запущенг")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
