package entity

import (
	"encoding/json"
	"fmt"
)

type ReportMethod string
type PeriodType string
type Period string

const (
	REPORT_MSFO       string     = "MSFO"
	REPORT_RSBU       string     = "RSBU"
	PERIOD_TYPE_YEAR  PeriodType = "YEAR"
	PERIOD_TYPE_MONTH PeriodType = "MONTH"
)

type FundamentalHeader struct {
	Ticker       string
	ReportMethod string
	PeriodType   string
	Period       string
	ReportUrl    string
}

type Fundamental struct {
	Date               string          `html:"date"`
	Currency           string          `html:"currency"`
	Revenue            FundamentalItem `html:"revenue"`
	OperatingIncome    FundamentalItem `html:"operating_income"`
	EBITDA             FundamentalItem `html:"ebitda"`
	NetIncome          FundamentalItem `html:"net_income"`
	Ocf                FundamentalItem `html:"ocf"`
	Capex              FundamentalItem `html:"capex"`
	Fcf                FundamentalItem `html:"fcf"`
	DividendPayout     FundamentalItem `html:"dividend_payout"`
	Dividend           FundamentalItem `html:"dividend"`
	DivYield           FundamentalItem `html:"div_yield"`
	DivPayoutRatio     FundamentalItem `html:"div_payout_ratio"`
	Opex               FundamentalItem `html:"opex"`
	Amortization       FundamentalItem `html:"amortization"`
	EmploymentExpenses FundamentalItem `html:"employment_expenses"`
	InterestExpenses   FundamentalItem `html:"interest_expenses"`
	Assets             FundamentalItem `html:"assets"`
	NetAssets          FundamentalItem `html:"net_assets"`
	Debt               FundamentalItem `html:"debt"`
	Cash               FundamentalItem `html:"cash"`
	NetDebt            FundamentalItem `html:"net_debt"`
	CommonShare        FundamentalItem `html:"common_share"`
	NumberOfShares     FundamentalItem `html:"number_of_shares"`
	MarketCap          FundamentalItem `html:"market_cap"`
	Ev                 FundamentalItem `html:"ev"`
	BookValue          FundamentalItem `html:"book_value"`
	Eps                FundamentalItem `html:"eps"`
	FcfShare           FundamentalItem `html:"fcf_share"`
	BvShare            FundamentalItem `html:"bv_share"`
	EbitdaMargin       FundamentalItem `html:"ebitda_margin"`
	NetMargin          FundamentalItem `html:"net_margin"`
	FcfYield           FundamentalItem `html:"fcf_yield"`
	Roe                FundamentalItem `html:"roe"`
	Roa                FundamentalItem `html:"roa"`
	PE                 string          `html:"p_e"`
	PFcf               FundamentalItem `html:"p_fcf"`
	PS                 string          `html:"p_s"`
	PBv                string          `html:"p_bv"`
	EvEbitda           FundamentalItem `html:"ev_ebitda"`
	DebtEbitda         FundamentalItem `html:"debt_ebitda"`
	RAndDCapex         FundamentalItem `html:"r_and_d_capex"`
	CapexRevenue       FundamentalItem `html:"capex_revenue"`
}

type FundamentalItem struct {
	Name    string
	Value   string
	Measure string
}

func SetFundamentalValue(f *Fundamental, htmlTag string, value interface{}, name interface{}, measure interface{}) {
	switch htmlTag {
	case "date":
		f.Date = value.(string)
	case "currency":
		f.Currency = value.(string)
	case "revenue":
		f.Revenue = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "operating_income":
		f.OperatingIncome = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "ebitda":
		f.EBITDA = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "net_income":
		f.NetIncome = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "ocf":
		f.Ocf = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "capex":
		f.Capex = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "fcf":
		f.Fcf = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "dividend_payout":
		f.DividendPayout = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "dividend":
		f.Dividend = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "div_yield":
		f.DivYield = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "div_payout_ratio":
		f.DivPayoutRatio = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "opex":
		f.Opex = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "amortization":
		f.Amortization = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "employment_expenses":
		f.EmploymentExpenses = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "interest_expenses":
		f.InterestExpenses = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "assets":
		f.Assets = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "net_assets":
		f.NetAssets = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "debt":
		f.Debt = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "cash":
		f.Cash = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "net_debt":
		f.NetDebt = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "common_share":
		f.CommonShare = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "number_of_shares":
		f.NumberOfShares = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "market_cap":
		f.MarketCap = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "ev":
		f.Ev = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "book_value":
		f.BookValue = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "eps":
		f.Eps = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "fcf_share":
		f.FcfShare = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "bv_share":
		f.BvShare = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "ebitda_margin":
		f.EbitdaMargin = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "net_margin":
		f.NetMargin = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "fcf_yield":
		f.FcfYield = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "roe":
		f.Roe = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "roa":
		f.Roa = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "p_e":
		f.PE = value.(string)
	case "p_fcf":
		f.PFcf = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "p_s":
		f.PS = value.(string)
	case "p_bv":
		f.PBv = value.(string)
	case "ev_ebitda":
		f.EvEbitda = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "debt_ebitda":
		f.DebtEbitda = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "r_and_d_capex":
		f.RAndDCapex = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	case "capex_revenue":
		f.CapexRevenue = FundamentalItem{Value: value.(string), Name: name.(string), Measure: measure.(string)}
	default:
		fmt.Println("Unknown HTML tag:", htmlTag)
	}
}

func FundamentalToJson(data Fundamental) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {

	}
	return jsonData
}
