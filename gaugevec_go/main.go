package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	C = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "count",
		Help: "count",
	}, []string{"Data01", "Data02", "Data03"})
)

func init() {
	// C.GetMetricWith()
	prometheus.MustRegister(C)
}

func main() {

	//goroutine處理工作

	r := http.NewServeMux()
	r.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", r)
}
