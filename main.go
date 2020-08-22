package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cputemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	hdFailures = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "Number of hard-disk errors.",
	}, []string{"device"})
)

func init() {
	prometheus.MustRegister(cputemp)
	prometheus.MustRegister(hdFailures)
}

func main() {

	cputemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9091", nil)

	fmt.Println("ok")
}
