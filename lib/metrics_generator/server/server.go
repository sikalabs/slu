package server

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sikalabs/slu/version"

	"gopkg.in/yaml.v3"
)

var promInfo = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "example_info",
	Help: "Instance info",
	ConstLabels: prometheus.Labels{
		"version": version.Version,
		"cmd0":    "slu",
		"cmd1":    "promdemo",
	},
})

var promRequestDurationSeconds = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: "example",
		Name:      "request_duration_seconds",
		Help:      "Request duration in seconds",
		Buckets:   []float64{.01, .05, .1, .2, .5, 1, 2, 5},
	},
	[]string{"status_code", "path", "method"},
)

type ServerStateMetrics struct {
	StatusCode            int     `yaml:"StatusCode"`
	Path                  string  `yaml:"Path"`
	Method                string  `yaml:"Method"`
	Duration              float64 `yaml:"Duration"`
	DurationDeviationPerc int     `yaml:"DurationDeviationPerc"`
	Rate                  int     `yaml:"Rate"`
	RateDeviationPerc     int     `yaml:"RateDeviationPerc"`
}

type ServerState struct {
	Metrics []ServerStateMetrics `yaml:"Metrics"`
}

type ServerConfig struct {
	Metrics []ServerStateMetrics `yaml:"Metrics"`
}

var State ServerState

func getSleepTime(rate int) time.Duration {
	return time.Duration(time.Duration(1000./float64(rate)) * time.Millisecond)
}

func addDeviationPerc(n float64, d int) float64 {
	if d == 0 {
		return n
	}
	r := float64(-d/2+rand.Intn(d)) / 100.
	return n + n*r
}

func runMetrics() {
	window := 60
	for {
		for _, metric := range State.Metrics {
			go func(m ServerStateMetrics) {
				rate := int(addDeviationPerc(float64(window*m.Rate), m.RateDeviationPerc))
				for i := 0; i < rate; i++ {
					promRequestDurationSeconds.WithLabelValues(
						strconv.Itoa(m.StatusCode),
						m.Path,
						m.Method,
					).Observe(addDeviationPerc(m.Duration, m.DurationDeviationPerc))
					time.Sleep(getSleepTime(m.Rate))
				}
			}(metric)
		}
		time.Sleep(time.Duration(window) * time.Second)
	}
}

func Server(addr string, config ServerConfig) {
	prometheus.MustRegister(promRequestDurationSeconds)
	prometheus.MustRegister(promInfo)
	promInfo.Set(1)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"_version": version.Version,
			"cmd0":     "slu",
			"cmd1":     "promdemo",
			"metrics":  "/metrics",
		})
	})

	State = ServerState{
		Metrics: config.Metrics,
	}

	go runMetrics()

	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[slu " + version.Version + "] Server started on 0.0.0.0:" + port + ", see http://127.0.0.1:" + port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func ServerWithConfig(addr, configPath string) {
	var config ServerConfig
	var err error
	f, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
	Server(addr, config)
}

func ServerWithDefaultConfig(addr string) {
	config := ServerConfig{
		Metrics: []ServerStateMetrics{
			{
				StatusCode:            200,
				Path:                  "/foo",
				Method:                "GET",
				Duration:              0.3,
				DurationDeviationPerc: 10,
				Rate:                  10,
				RateDeviationPerc:     20,
			},
			{
				StatusCode:            200,
				Path:                  "/bar",
				Method:                "GET",
				Duration:              0.1,
				DurationDeviationPerc: 10,
				Rate:                  20,
				RateDeviationPerc:     40,
			},
			{
				StatusCode:            200,
				Path:                  "/baz",
				Method:                "GET",
				Duration:              0.7,
				DurationDeviationPerc: 10,
				Rate:                  11,
				RateDeviationPerc:     10,
			},
			{
				StatusCode:            502,
				Path:                  "/baz",
				Method:                "GET",
				Duration:              0.01,
				DurationDeviationPerc: 10,
				Rate:                  1,
				RateDeviationPerc:     10,
			},
		},
	}
	Server(addr, config)
}
