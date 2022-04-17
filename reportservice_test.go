package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCreateReport(t *testing.T) {
	var tests = []struct {
		Name      string
		Portfolio Portfolio
		Want      PortfolioReport
		MockFunc  func()
	}{
		{
			Name: "Happy Path",
			Portfolio: Portfolio{
				Holdings: []Holding{
					{
						Ticker:       "AZN",
						PurchaseDate: time.Now(),
						Quantity:     6,
						BoughtAt:     461.31,
						Provider:     "LSE",
					},
				},
			},
			MockFunc: func() {
				GetCloseForDate = func(ticker string, date time.Time, provider string) float64 {
					return 10500
				}
			},
			Want: PortfolioReport{
				StockReports: []StockReport{
					{
						Ticker:             "AZN",
						UnitsHeld:          6,
						Price:              10500.00,
						Value:              630.00,
						Cost:               461.31,
						PercentageGainLoss: 36.57,
						ValueGainLoss:      168.69,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		testName := tt.Name
		t.Run(testName, func(t *testing.T) {
			tt.MockFunc()
			ans := createReport(tt.Portfolio)
			if !reflect.DeepEqual(ans, tt.Want) {
				t.Errorf("got %v, want %v", ans, tt.Want)
			}
		})
	}
}
