package scraper

import (
	"errors"
	"fin_fundamentals/internal/entity"
	"fin_fundamentals/internal/log"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"reflect"
	"strings"
)

func GetSmartLabUri(url string, ticker string, repMethod string) string {
	return fmt.Sprintf("%s%s/f/q/%s/", url, ticker, repMethod)
}

func ScrapSmartLabSecurity(uri string, ticker string, reportMethod string) map[entity.FundamentalHeader]entity.Fundamental {
	res, err := http.Get(uri)
	if err != nil {
		log.Error("Ошбика http запроса:", err)
	}
	defer res.Body.Close()

	//file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//
	//log.SetOutput(file)

	if res.StatusCode != 200 {
		log.Error("status code error", errors.New("status code is not 200"))
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Error("Ошибка парсинга html", err)
	}

	if doc.Find("table.financials").Length() == 0 {
		log.Error("данные для парсинга не найдены на странице", errors.New(""))
	}

	var fundamentals map[entity.FundamentalHeader]entity.Fundamental = make(map[entity.FundamentalHeader]entity.Fundamental)

	doc.Find("table.financials").Each(func(i int, table *goquery.Selection) {
		table.Find("tr.header_row td > strong").Each(func(i int, col *goquery.Selection) {
			colEntry := i + 1

			col.Each(func(i int, headerCol *goquery.Selection) {
				fundamentalToSet := entity.Fundamental{}

				reportTagParent := table.Find(`tr[field="report_url"] > td`).Eq(colEntry)
				reportUrl, _ := reportTagParent.Find(`a`).Attr("href")

				// Идет перебор атрибутов для парсинга в рамках столбца
				t := reflect.TypeOf(entity.Fundamental{})
				for i := 0; i < t.NumField(); i++ {
					field := t.Field(i)
					html := field.Tag.Get("html") // Получаем значение тега html

					val := strings.TrimSpace(table.Find(fmt.Sprintf(`tr[field="%s"] > td`, html)).Eq(colEntry).Text())
					name := strings.TrimSpace(table.Find(fmt.Sprintf(`tr[field="%s"] > th > a`, html)).Eq(0).Text())
					measure := strings.TrimSpace(table.Find(fmt.Sprintf(`tr[field="%s"] > th > span`, html)).Eq(0).Text())

					err := entity.SetFundamentalValue(&fundamentalToSet, html, val, name, measure)
					if err != nil {
						log.Error("Ошибка при наполнении сущности", err)
					}
				}

				header := entity.FundamentalHeader{
					Ticker:       ticker,
					Period:       "quarter",
					ReportMethod: reportMethod,
					Report:       headerCol.Text(),
					ReportUrl:    reportUrl,
					SourceUrl:    uri,
				}

				fundamentals[header] = fundamentalToSet
			})
		})

	})

	return fundamentals
}
