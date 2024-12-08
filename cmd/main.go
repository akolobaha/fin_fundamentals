package main

import (
	"context"
	"fin_fundamentals/cmd/commands"
	"fin_fundamentals/internal/config"
	"fin_fundamentals/internal/transport"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultEnvFilePath = ".env"

func main() {
	cfg, err := config.Parse(defaultEnvFilePath)
	if err != nil {
		panic("Ошибка парсинга конфигов")
	}

	rabbit := transport.New()
	rabbit.InitConn(cfg)
	defer rabbit.ConnClose()
	rabbit.DeclareQueue(cfg.RabbitQueue)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-exit
		cancel()
	}()

	timeTicker := time.NewTicker(100000)
	defer timeTicker.Stop()

	slog.Info("Сервис сбора отчетности запущен")
	commands.RunParser(ctx, cfg, rabbit)

	for {
		select {
		case <-cfg.TickInterval.C:
			commands.RunParser(ctx, cfg, rabbit)
		case <-ctx.Done():
			rabbit.ConnClose()
			timeTicker.Stop()
			slog.Info("Сбор фундаментальных данных остановлен")
			return
		}
	}

}
