package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func TrackIp(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"ip": r.RemoteAddr,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	port := 8080
	http.HandleFunc("/ping", TrackIp)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
