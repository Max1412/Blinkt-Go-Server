package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Max1412/blinkt_server/internal/app/server_backend"
)

func main() {
	defer server_backend.LedCleaner()

	fmt.Println("Starting server...")

	http.HandleFunc("/", server_backend.Handler)

	http.HandleFunc("/SolidColor/", server_backend.HandlerLEDSolidColor)
	http.HandleFunc("/Progress/", server_backend.HandlerLEDProgress)
	http.HandleFunc("/WakeUp/", server_backend.HandlerLEDWakeUp)

	http.HandleFunc("/stop/", server_backend.HandlerStopAsync)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
