package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func serveRandomFile(w http.ResponseWriter, r *http.Request) {
	rnd := rand.Intn(10)
	if rnd < 5 {
		http.ServeFile(w, r, "./dummy.png")
        log.Printf("Received request for dummy.png")
		return
	}
	if rnd < 9 {
		http.ServeFile(w, r, "./dummy.pdf")
        log.Printf("Received request for dummy.pdf")
		return
	}
	http.ServeFile(w, r, "./corrupt-dummy.pdf")
    log.Printf("Received request for corrupt-dummy.pdf")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", serveRandomFile)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/readiness", readinessHandler)
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal(err)
	}
}
