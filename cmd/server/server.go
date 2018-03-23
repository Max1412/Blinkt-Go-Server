package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Max1412/blinkt_server/internal/app/serverbackend"
)

func main() {
	defer serverbackend.LedCleaner()

	fmt.Println("Starting server...")

	http.HandleFunc("/", serverbackend.Handler)

	http.HandleFunc("/SolidColor/", serverbackend.HandlerLEDSolidColor)
	http.HandleFunc("/Progress/", serverbackend.HandlerLEDProgress)
	http.HandleFunc("/WakeUp/", serverbackend.HandlerLEDWakeUp)

	http.HandleFunc("/stop/", serverbackend.HandlerStopAsync)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
