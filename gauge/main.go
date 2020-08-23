package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	Count = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total",
		Help: "total count",
	})
)

func init() {
	prometheus.MustRegister(Count)
}

func main() {

	Count.Set(100)

	r := http.NewServeMux()
	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/", index)
	r.HandleFunc("/Inc", Inc)
	r.HandleFunc("/Dec", Dec)
	r.HandleFunc("/Add", Add)
	r.HandleFunc("/Sub", Sub)
	http.ListenAndServe(":8080", r)

	fmt.Println("OK")
}

func index(w http.ResponseWriter, r *http.Request) {
	// Count.SetToCurrentTime()
	fmt.Fprintln(w, "ok")
}

func Inc(w http.ResponseWriter, r *http.Request) {
	Count.Inc()
}

func Dec(w http.ResponseWriter, r *http.Request) {
	Count.Dec()
}

func Add(w http.ResponseWriter, r *http.Request) {
	c := r.URL.Query()

	c_str := c.Get("c")

	count, err := strconv.Atoi(c_str)

	if err != nil {
		fmt.Println(err)
		return
	}

	Count.Add(float64(count))
}

func Sub(w http.ResponseWriter, r *http.Request) {
	c := r.URL.Query()

	c_str := c.Get("c")

	count, err := strconv.Atoi(c_str)

	if err != nil {
		fmt.Println(err)
		return
	}

	Count.Sub(float64(count))
}
