package main

import (
	"time"
)

func createReport(portfolio Portfolio) PortfolioReport {
	portfolioReport := PortfolioReport{}

	for _, holding := range portfolio.Holdings {
		var currentPrice float64

		hp := getHistoricalPricesForTicker(holding.Provider)
		if holding.Provider == "LSE" {
			currentPrice = getLsePrice(holding.Ticker)
		} else {
			currentPrice = getCloseForDate(hp, time.Now()).Close * exchangeRate * 100
		}

		percentGainLoss := roundTo(10, 2)
		valueGainLoss := roundTo((currentPrice*holding.Quantity/100)-(holding.BoughtAt), 1)

		sp := StockReport{
			Ticker:             holding.Ticker,
			UnitsHeld:          holding.Quantity,
			Price:              roundTo(currentPrice, 2),
			Value:              roundTo(currentPrice*holding.Quantity/100, 2),
			Cost:               roundTo(holding.BoughtAt, 2),
			PercentageGainLoss: percentGainLoss,
			ValueGainLoss:      valueGainLoss,
		}
		portfolioReport.StockReports = append(portfolioReport.StockReports, sp)
	}
	return portfolioReport
}
