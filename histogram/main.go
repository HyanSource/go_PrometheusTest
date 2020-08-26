package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Count = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "total",
		Help:    "total",
		Buckets: prometheus.LinearBuckets(20, 5, 5), //初始20 每個區間差距5 共有5個區間
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
