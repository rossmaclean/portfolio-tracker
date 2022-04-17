package main

import (
	"time"
)

var fetchCost = false

func createReport(portfolio Portfolio) PortfolioReport {
	portfolioReport := PortfolioReport{}

	for _, holding := range portfolio.Holdings {
		var currentPrice = GetCloseForDate(holding.Ticker, time.Now(), holding.Provider)
		//var boughtAt = holding.BoughtAt
		//if holding.Provider == "YAHOO" {
		//	boughtAt = GetCloseForDate(holding.Ticker, holding.PurchaseDate, holding.Provider)
		//}

		//log.Printf("fetched boughtat: %f, supplied boughtat: %f", boughtAt, holding.BoughtAt)

		value := currentPrice * holding.Quantity / 100
		cost := holding.Cost / 100
		//boughtAt := holding.Cost / holding.Quantity

		valueGainLoss := value - cost
		percentGainLoss := (value - cost) / cost * 100

		sp := StockReport{
			Ticker:             holding.Ticker,
			UnitsHeld:          holding.Quantity,
			Price:              roundTo(currentPrice, 2),
			Value:              roundTo(value, 2),
			Cost:               roundTo(cost, 2),
			PercentageGainLoss: roundTo(percentGainLoss, 2),
			ValueGainLoss:      roundTo(valueGainLoss, 2),
		}
		portfolioReport.StockReports = append(portfolioReport.StockReports, sp)
	}
	return portfolioReport
}
