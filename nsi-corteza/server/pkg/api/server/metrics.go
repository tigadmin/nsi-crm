package server

import (
    "net/http"

    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// MetricsMiddleware is the request logger that provides metrics to prometheus
func metricsMiddleware(name string) func(http.Handler) http.Handler {
	return chiprometheus.NewMiddleware(name)
}

func metricsMount(r chi.Router, username, password string) {
	r.Route("/metrics", func(r chi.Router) {
		r.Use(basicauth.New("Metrics", map[string][]string{
			username: {password},
		}))
		r.Handle("/", promhttp.Handler())
	})
}
