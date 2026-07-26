package main

import (
	"log"
	"net/http"

	"github.com/namwamba/task-api/internal/router"
)

func main() {
	r := router.New()

	log.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
