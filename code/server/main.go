package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Ok bool `json:"ok"`
}

func TrackIp(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Ok: true,
	}

	fmt.Println(response)
	json.NewEncoder(w).Encode(response)
}

func main() {
	port := 8080
	http.HandleFunc("/ping", TrackIp)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
