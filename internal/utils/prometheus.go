package utils

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	RequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Latency of HTTP requests",
			Buckets: prometheus.DefBuckets, // Default buckets: [0.005, 0.01, 0.025, ...]
		},
		[]string{"method", "path", "status"},
	)
)

func InitPrometheus() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(RequestLatency)
}
