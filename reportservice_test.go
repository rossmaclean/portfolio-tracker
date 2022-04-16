package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/svarlamov/goyhfin"
)

func TestCreateReport(t *testing.T) {
	var tests = []struct {
		Name      string
		portfolio Portfolio
		want      PortfolioReport
		mockFunc  func()
	}{
		{
			Name: "Happy Path",
			portfolio: Portfolio{
				Holdings: []Holding{
					Holding{
						Ticker:       "AZN",
						PurchaseDate: time.Now(),
						Quantity:     6,
						BoughtAt:     461.31,
						Provider:     "LSE",
					},
				},
			},
			mockFunc: func() {
				getCloseForDate = func(quaotes []goyhfin.Quote, date time.Time) goyhfin.Quote {
					return goyhfin.Quote{
						OpensAt: time.Now(),	
						Open: 0,
						High: 0,
						Low: 0,
						Close: 10500,
						Volume: 1,
						ClosesAt: time.Now(),
						Period: time.Now().Sub(time.Now()), // This is bullshit, fix it
					}
				}

				getHistoricalPricesForTicker = func(ticker string) []goyhfin.Quote {
					return []goyhfin.Quote{}
				}

				getLsePrice = func(ticker string) float64 {
					return 10500
				}

			},
		}, {
			want: PortfolioReport{
				StockReports: []StockReport{
					StockReport{
						Ticker:             "AZN",
						UnitsHeld:          6,
						Price:              10500.00,
						Value:              630.00,
						Cost:               461.31,
						PercentageGainLoss: 36.57,
						ValueGainLoss:      168.69,
					},
//					StockReport{
//						Ticker:             "IAG",
//						UnitsHeld:          534,
//						Price:              140.36,
//						Value:              749.52,
//						Cost:               511.35,
//						PercentageGainLoss: 46.58,
//						ValueGainLoss:      238.17,
//					},
				},
			},
		},
	}

	for _, tt := range tests {
		testname := tt.Name
		t.Run(testname, func(t *testing.T) {
			tt.mockFunc()
			ans := createReport(tt.portfolio)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
