package commands

import (
	"fin_fundamentals/internal/entity"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func GetSmartLabUri(url string, ticker string, repMethod entity.ReportMethod) string {
	return fmt.Sprintf("%s%s/f/q/%s/", url, ticker, repMethod)
}

func ScrapSmartLabSecurity(uri string) map[string]entity.Fundamental {
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var fundamentals map[string]entity.Fundamental = make(map[string]entity.Fundamental)

	doc.Find("table.financials").Each(func(i int, table *goquery.Selection) {
		table.Find("tr.header_row td > strong").Each(func(i int, col *goquery.Selection) {
			colEntry := i + 1

			col.Each(func(i int, headerCol *goquery.Selection) {
				fundamentalToSet := entity.Fundamental{}
				fundamentals[headerCol.Text()] = fundamentalToSet

				// Идет перебор атрибутов для парсинга в рамках столбца
				t := reflect.TypeOf(fundamentals[headerCol.Text()])

				for i := 0; i < t.NumField(); i++ {
					field := t.Field(i)
					html := field.Tag.Get("html") // Получаем значение тега html

					val := strings.TrimSpace(table.Find(fmt.Sprintf(`tr[field="%s"] > td`, html)).Eq(colEntry).Text())
					name := strings.TrimSpace(table.Find(fmt.Sprintf(`tr[field="%s"] > th > a`, html)).Eq(0).Text())
					measure := strings.TrimSpace(table.Find(fmt.Sprintf(`tr[field="%s"] > th > span`, html)).Eq(0).Text())

					entity.SetFundamentalValue(&fundamentalToSet, html, val, name, measure)
				}

				fundamentals[headerCol.Text()] = fundamentalToSet
			})
		})

	})

	return fundamentals
}
