package main

import (
	"context"
	"fin_fundamentals/cmd/commands"
	"fin_fundamentals/internal/config"
	"fin_fundamentals/internal/log"
	"fin_fundamentals/internal/monitoring"
	"fin_fundamentals/internal/transport"
	"os"
	"os/signal"
	"syscall"
)

const defaultEnvFilePath = ".env"

func init() {
	monitoring.RegisterPrometheus()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cfg, err := config.Parse(defaultEnvFilePath)
	if err != nil {
		panic("Ошибка парсинга конфигов")
	}

	monitoring.RunPrometheusServer(cfg.GetPrometheusURL())
	defer monitoring.StopPrometheusServer(ctx)

	rabbit := transport.New()
	rabbit.InitConn(cfg)
	defer rabbit.ConnClose()
	rabbit.DeclareQueue(cfg.RabbitQueue)

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-exit
		cancel()
	}()

	log.Info("Сервис сбора отчетности запущен")
	commands.RunParser(ctx, cfg, rabbit)
	for {
		select {
		case <-cfg.TickInterval.C:
			commands.RunParser(ctx, cfg, rabbit)
		case <-ctx.Done():
			log.Info("Сервис сбора отчетности остановлен")
			return
		}
	}

}
