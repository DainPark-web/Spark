package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on port 8080")
	port := 8000
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
