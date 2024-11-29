package main

import (
	"fin_fundamentals/cmd/commands"
	"fin_fundamentals/internal/config"
	"fin_fundamentals/internal/entity"
	"fin_fundamentals/internal/transport"
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

	for _, ticker := range entity.Tickers {
		for _, method := range []string{entity.REPORT_RSBU, entity.REPORT_MSFO} {
			uri := commands.GetSmartLabUri(cfg.SourceUrl, ticker, method)
			data := commands.ScrapSmartLabSecurity(uri, ticker, method)

			for header, item := range data {
				jsonData := entity.FundamentalToJson(item)

				rabbit.SendMsg(jsonData, header)
			}
		}
	}

}
