package monitoring

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"log/slog"
	"net/http"
)

var server *http.Server

var (
	FundamentalErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "parsing_errors_total",
			Help: "Parsing error",
		},
		[]string{"message"},
	)
	FundamentalSuccessCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "parsing_success_total",
			Help: "Parsing success",
		},
		[]string{"message"},
	)
)

func RegisterPrometheus() {
	prometheus.MustRegister(FundamentalErrorCount)
	prometheus.MustRegister(FundamentalSuccessCount)
}

func RunPrometheusServer(url string) {
	server = &http.Server{
		Addr: url,
	}

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start prometheus server: %v", err)
		}
	}()
}

func StopPrometheusServer(ctx context.Context) {
	if server != nil {
		if err := server.Shutdown(ctx); err != nil {
			slog.Error("Server forced to shutdown: %v", "error", err)
		}
		slog.Info("Server exited gracefully")
	} else {
		slog.Warn("Server is not running")
	}
}
