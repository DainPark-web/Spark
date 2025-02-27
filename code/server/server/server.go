package server

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

func StartServer() {
	port := 8080
	port2 := 8081

	// ServeMux1
	mux := http.NewServeMux()

	mux.HandleFunc("/pings", TrackIp)

	// Server1
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
	}()

	// Server2
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port2), mux))
	}()

	select {}
}
