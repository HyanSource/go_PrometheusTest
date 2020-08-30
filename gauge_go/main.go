package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//需要統計的資料
var (
	Data01 = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "Data01",
		Help: "Data01 total",
	})
	Data02 = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "Data02",
		Help: "Data02 total",
	})
	Data03 = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "Data03",
		Help: "Data03 total",
	})
)

//初始化
func init() {
	prometheus.MustRegister(Data01)
	prometheus.MustRegister(Data02)
	prometheus.MustRegister(Data03)

	Data01.Set(0)
	Data02.Set(0)
	Data03.Set(0)
}

func main() {

	//Data01
	go func() {
		for {
			t := rand.Int31n(2)
			c := float64(rand.Int31n(50))

			switch t {
			case 0:
				Data01.Add(c)
				break
			case 1:
				Data01.Sub(c)
				break
			}

			time.Sleep(5 * time.Second)
		}
	}()

	//Data02
	go func() {
		for {
			t := rand.Int31n(2)
			c := float64(rand.Int31n(10))

			switch t {
			case 0:
				Data02.Add(c)
				break
			case 1:
				Data02.Sub(c)
				break
			}

			time.Sleep(5 * time.Second)
		}
	}()

	//Data03
	go func() {
		t := rand.Int31n(2)
		switch t {
		case 0:
			Data03.Inc()
			break
		case 1:
			Data03.Dec()
			break
		}
		time.Sleep(5 * time.Second)
	}()

	fmt.Println("OK")

	r := http.NewServeMux()
	r.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", r)
}
