package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	version = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "version",
		Help: "Version information about this binary",
		ConstLabels: map[string]string{
			"version": "v0.1.0",
		},
	})
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"code", "method"})
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {

	r := prometheus.NewRegistry()
	r.MustRegister(httpRequestsTotal)
	r.MustRegister(version)

	rand.Seed(time.Now().UnixNano())

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/readiness", readinessHandler)

	http.Handle("/", promhttp.InstrumentHandlerCounter(httpRequestsTotal, handler))
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))

	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal(err)
	}
}
