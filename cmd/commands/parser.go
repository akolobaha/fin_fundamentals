package commands

import (
	"fin_fundamentals/internal/config"
	"fin_fundamentals/internal/entity"
	"fin_fundamentals/internal/scraper"
	"fin_fundamentals/internal/transport"
)

func RunParser(cfg *config.Config, rabbit *transport.Rabbitmq) {
	for _, ticker := range entity.Tickers {
		for _, method := range []string{entity.REPORT_RSBU, entity.REPORT_MSFO} {
			uri := scraper.GetSmartLabUri(cfg.SourceUrl, ticker, method)
			data := scraper.ScrapSmartLabSecurity(uri, ticker, method)

			for header, item := range data {
				jsonData := entity.FundamentalToJson(item)
				rabbit.SendMsg(jsonData, header)
			}
		}
	}
}
