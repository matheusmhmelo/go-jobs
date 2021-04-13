package main

import (
	"github.com/matheusmhmelo/go-jobs/internal"
	"github.com/matheusmhmelo/go-jobs/internal/repository"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	repository.Start()

	h := internal.NewServer()

	log.Println("running server in port " + port)
	log.Fatal(http.ListenAndServe(":"+port, h))
}
