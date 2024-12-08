package commands

import (
	"context"
	"fin_fundamentals/internal/config"
	"fin_fundamentals/internal/entity"
	"fin_fundamentals/internal/scraper"
	"fin_fundamentals/internal/transport"
	"fmt"
)

func RunParser(ctx context.Context, cfg *config.Config, rabbit *transport.Rabbitmq) {
	for _, ticker := range entity.Tickers {
		for _, method := range []string{entity.REPORT_RSBU, entity.REPORT_MSFO} {
			select {
			case <-ctx.Done():
				// Обработка отмены
				fmt.Println("Parsing was cancelled")
				return
			default:
				uri := scraper.GetSmartLabUri(cfg.SourceUrl, ticker, method)
				data := scraper.ScrapSmartLabSecurity(uri, ticker, method)

				for header, item := range data {
					jsonData := entity.FundamentalToJson(item)
					rabbit.SendMsg(jsonData, header)
				}
			}
		}
	}
}
