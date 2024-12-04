package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ReportMethod string
type Report string
type Period string

const (
	REPORT_MSFO string = "MSFO"
	REPORT_RSBU string = "RSBU"
)

type FundamentalHeader struct {
	Ticker       string
	ReportMethod string
	Report       string
	Period       string
	ReportUrl    string
	SourceUrl    string
}

type Fundamental struct {
	Date      string `html:"date"`
	Currency  string `html:"currency"`
	Revenue   uint64 `html:"revenue"`
	NetIncome int64  `html:"net_income"`
	BookValue uint64 `html:"book_value"`
}

type FundamentalItem struct {
	Name    string
	Value   string
	Measure string
}

func SetFundamentalValue(f *Fundamental, htmlTag string, value interface{}, name interface{}, measure interface{}) error {
	treamedValue := strings.ReplaceAll(value.(string), " ", "")
	switch htmlTag {
	case "date":
		f.Date = treamedValue
	case "currency":
		f.Currency = treamedValue
	case "revenue":
		val, err := convertStringBillionToInt(treamedValue, measure.(string))
		if err != nil {
			return err
		}
		f.Revenue = val
	case "net_income":
		val, err := convertStringBillionToInt64(treamedValue, measure.(string))
		if err != nil {
			return err
		}
		f.NetIncome = val
	case "book_value":
		val, err := convertStringBillionToInt(treamedValue, measure.(string))
		if err != nil {
			return err
		}
		f.BookValue = val
	default:
		fmt.Println("Unknown HTML tag:", htmlTag)
	}

	return nil
}

// млн
func convertStringMillionToInt(value string, measure string) (uint64, error) {
	var result float64
	if len(value) == 0 {
		return 0, errors.New("shares number is empty")
	}

	result, _ = strconv.ParseFloat(value, 64)

	if measure == "млн" {
		result *= 1000000
	}

	return uint64(result), nil
}

func convertStringBillionToInt64(value string, measure string) (int64, error) {
	var result float64
	if len(value) == 0 {
		return 0, errors.New("shares number is empty")
	}

	result, _ = strconv.ParseFloat(value, 64)

	if measure == "млрд руб" {
		result *= 1000000000
	}

	return int64(result), nil
}

func convertStringBillionToInt(value string, measure string) (uint64, error) {
	var result float64
	if len(value) == 0 {
		return 0, errors.New("shares number is empty")
	}

	result, _ = strconv.ParseFloat(value, 64)

	if measure == "млрд руб" {
		result *= 1000000000
	}

	return uint64(result), nil
}

func FundamentalToJson(data Fundamental) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {

	}
	return jsonData
}
