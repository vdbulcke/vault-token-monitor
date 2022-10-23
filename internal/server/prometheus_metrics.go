package server

import "github.com/prometheus/client_golang/prometheus"

var (
	promMetricTokenExpirationSeconds = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "vault_monitor",
			Name:      "token_expiration_timestamp_seconds",
			Help:      "Expiration Date of Token as Unix Timestamp in seconds",
		},
		[]string{
			"display_name",
			"accessor",
			"auto_renew",
		},
	)

	promMetricTokenErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "vault_monitor",
			Name:      "token_error_count",
			Help:      "Number of error while processing the token",
		},
		[]string{
			"accessor",
		},
	)

	promMetricTokenRenewCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "vault_monitor",
			Name:      "token_renew_count",
			Help:      "Number of renew for the token",
		},
		[]string{
			"display_name",
			"accessor",
		},
	)
)

// PrometheusMetricsRegister Register metrics with prometheus
func PrometheusMetricsRegister() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(promMetricTokenExpirationSeconds)
	prometheus.MustRegister(promMetricTokenErrorCount)
	prometheus.MustRegister(promMetricTokenRenewCount)

}
