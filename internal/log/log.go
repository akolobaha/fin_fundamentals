package log

import (
	"fin_fundamentals/internal/monitoring"
	"fmt"
	"log/slog"
)

func Error(additionalMessage string, err error) {
	if err != nil {
		msg := fmt.Sprintf("%s: %s", additionalMessage, err.Error())
		monitoring.FundamentalErrorCount.WithLabelValues(msg).Inc()
		slog.Error(msg)
	}
}

func Info(message string) {
	monitoring.FundamentalSuccessCount.WithLabelValues(message).Inc()
	slog.Info(message)
}
