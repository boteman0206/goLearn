package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "pw_requests_total",
		Help: "The total number of requests",
	})
	requestsFailed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "pw_requests_failed_total",
		Help: "The total number of failed requests",
	})
)

func init() {
	// Register the counters with Prometheus's default registry.
	prometheus.MustRegister(requestsTotal, requestsFailed)
}

func main() {
	// Create a new router.
	router := mux.NewRouter()

	// Handle requests to the / endpoint.
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Increment the total number of requests.
		requestsTotal.Inc()

		// Simulate a failure if a query parameter called "fail" is present.
		fmt.Println(r.URL.Query())
		if _, ok := r.URL.Query()["fail"]; ok {
			requestsFailed.Inc()
			http.Error(w, "Failed", http.StatusInternalServerError)
			return
		}

		// Send a response.
		w.Write([]byte("OK"))
	})

	// Expose the registered metrics via HTTP.
	router.Handle("/metrics", promhttp.Handler())

	// Start the HTTP server.
	log.Fatal(http.ListenAndServe(":8080", router))
}
