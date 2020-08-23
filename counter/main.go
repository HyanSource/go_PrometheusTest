package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Count = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "count",
		Help: "http count",
	})
)

func init() {
	prometheus.MustRegister(Count)
}

func main() {
	r := http.NewServeMux()
	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/", Inc)
	r.HandleFunc("/Add", Add)

	http.ListenAndServe(":8080", r)

	fmt.Println("OK")
}

func Inc(w http.ResponseWriter, r *http.Request) {
	Count.Inc()
}

func Add(w http.ResponseWriter, r *http.Request) {
	Count.Add(10)
}
