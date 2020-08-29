package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Count = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "",
		Help:       "",
		Objectives: map[float64]float64{},
	})
)

func main() {
	r := http.NewServeMux()
	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/", index)

	http.ListenAndServe(":8080", r)
	fmt.Println("OK")
}

func index(w http.ResponseWriter, r *http.Request) {
	Count.Observe(float64(rand.Int31n(50)))
}
