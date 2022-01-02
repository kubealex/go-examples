package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var reqCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "custom_request_count",
		Help: "Number of request handled by hello handler",
	},
)

func hello(w http.ResponseWriter, req *http.Request) {
	reqCounter.Inc()
	fmt.Fprintf(w, "Hello from this go sample")
}

func main() {
	prometheus.MustRegister(reqCounter)

	http.HandleFunc("/hello", hello)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8090", nil)
}
