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
	rabbit.InitConn()
	defer rabbit.ConnClose()
	rabbit.DeclareQueue("tickers")

	for _, ticker := range entity.Tickers {
		uri := commands.GetSmartLabUri(cfg.SourceUrl, ticker, entity.REPORT_MSFO)
		data := commands.ScrapSmartLabSecurity(uri)

		jsonData := entity.FundamentalToJson(data)
		rabbit.SendMsg(jsonData)
	}

}
