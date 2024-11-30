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
	Date     string `html:"date"`
	Currency string `html:"currency"`
	//Revenue            FundamentalItem `html:"revenue"`
	//OperatingIncome    FundamentalItem `html:"operating_income"`
	//EBITDA             FundamentalItem `html:"ebitda"`
	NetIncome uint64 `html:"net_income"`
	//Ocf                FundamentalItem `html:"ocf"`
	//Capex              FundamentalItem `html:"capex"`
	//Fcf                FundamentalItem `html:"fcf"`
	//DividendPayout     FundamentalItem `html:"dividend_payout"`
	Dividend float64 `html:"dividend"`
	//DivYield           FundamentalItem `html:"div_yield"`
	//DivPayoutRatio     FundamentalItem `html:"div_payout_ratio"`
	//Opex               FundamentalItem `html:"opex"`
	//Amortization       FundamentalItem `html:"amortization"`
	//EmploymentExpenses FundamentalItem `html:"employment_expenses"`
	//InterestExpenses   FundamentalItem `html:"interest_expenses"`
	//Assets             FundamentalItem `html:"assets"`
	//NetAssets          FundamentalItem `html:"net_assets"`
	//Debt               FundamentalItem `html:"debt"`
	//Cash               FundamentalItem `html:"cash"`
	//NetDebt            FundamentalItem `html:"net_debt"`
	//CommonShare        FundamentalItem `html:"common_share"`
	NumberOfShares     uint64 `html:"number_of_shares"`
	NumberOfPrivShares uint64 `html:"number_of_priv_shares"`
	//MarketCap          FundamentalItem `html:"market_cap"`
	//Ev                 FundamentalItem `html:"ev"`
	BookValue uint64 `html:"book_value"`
	//Eps                FundamentalItem `html:"eps"`
	//FcfShare           FundamentalItem `html:"fcf_share"`
	//BvShare            FundamentalItem `html:"bv_share"`
	//EbitdaMargin       FundamentalItem `html:"ebitda_margin"`
	//NetMargin          FundamentalItem `html:"net_margin"`
	//FcfYield           FundamentalItem `html:"fcf_yield"`
	//Roe                FundamentalItem `html:"roe"`
	//Roa                FundamentalItem `html:"roa"`
	PE float64 `html:"p_e"`
	//PFcf               FundamentalItem `html:"p_fcf"`
	//PS                 string          `html:"p_s"`
	PBv float64 `html:"p_bv"`
	//EvEbitda           FundamentalItem `html:"ev_ebitda"`
	//DebtEbitda         FundamentalItem `html:"debt_ebitda"`
	//RAndDCapex         FundamentalItem `html:"r_and_d_capex"`
	//CapexRevenue       FundamentalItem `html:"capex_revenue"`
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
	//case "revenue":
	//	f.Revenue = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "operating_income":
	//	f.OperatingIncome = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "ebitda":
	//	f.EBITDA = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	case "net_income":
		val, err := convertStringBillionToInt(treamedValue, measure.(string))
		if err != nil {
			return err
		}
		f.NetIncome = val
	//case "ocf":
	//	f.Ocf = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "capex":
	//	f.Capex = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "fcf":
	//	f.Fcf = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "dividend_payout":
	//	f.DividendPayout = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	case "dividend":
		if treamedValue != "" {
			valueToSet, err := strconv.ParseFloat(treamedValue, 64)
			if err != nil {
				return err
			}
			f.Dividend = valueToSet
		}

	//case "div_yield":
	//	f.DivYield = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "div_payout_ratio":
	//	f.DivPayoutRatio = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "opex":
	//	f.Opex = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "amortization":
	//	f.Amortization = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "employment_expenses":
	//	f.EmploymentExpenses = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "interest_expenses":
	//	f.InterestExpenses = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "assets":
	//	f.Assets = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "net_assets":
	//	f.NetAssets = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "debt":
	//	f.Debt = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "cash":
	//	f.Cash = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "net_debt":
	//	f.NetDebt = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "common_share":
	//	f.CommonShare = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	case "number_of_shares":
		val, err := convertStringMillionToInt(treamedValue, measure.(string))
		if err != nil {
			return err
		}
		f.NumberOfShares = val
	case "number_of_priv_shares":
		val, err := convertStringMillionToInt(treamedValue, measure.(string))
		if err != nil {
			return err
		}
		f.NumberOfPrivShares = val
		//case "market_cap":
		//	f.MarketCap = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
		//case "ev":
		//	f.Ev = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	case "book_value":
		val, err := convertStringBillionToInt(treamedValue, measure.(string))
		if err != nil {
			return err
		}
		f.BookValue = val
	//case "eps":
	//	f.Eps = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "fcf_share":
	//	f.FcfShare = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "bv_share":
	//	f.BvShare = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "ebitda_margin":
	//	f.EbitdaMargin = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "net_margin":
	//	f.NetMargin = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "fcf_yield":
	//	f.FcfYield = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "roe":
	//	f.Roe = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "roa":
	//	f.Roa = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	case "p_e":
		if treamedValue != "" {
			valueToSet, err := strconv.ParseFloat(treamedValue, 64)
			if err != nil {
				return err
			}
			f.PE = valueToSet
		}

	//case "p_fcf":
	//	f.PFcf = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "p_s":
	//	f.PS = treamedValue
	case "p_bv":
		if treamedValue != "" {
			valueToSet, err := strconv.ParseFloat(treamedValue, 64)
			if err != nil {
				return err
			}
			f.PBv = valueToSet
		}

	//case "ev_ebitda":
	//	f.EvEbitda = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "debt_ebitda":
	//	f.DebtEbitda = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "r_and_d_capex":
	//	f.RAndDCapex = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
	//case "capex_revenue":
	//	f.CapexRevenue = FundamentalItem{Value: treamedValue, Name: name.(string), Measure: measure.(string)}
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
