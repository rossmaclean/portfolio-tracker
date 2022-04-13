package main

import "time"

type StockReport struct {
	Ticker             string
	UnitsHeld          float64
	Price              float64
	Value              float64
	Cost               float64
	PercentageGainLoss float64
	ValueGainLoss      float64
}

type PortfolioReport struct {
	StockReports []StockReport
}

type Holding struct {
	Ticker       string
	PurchaseDate time.Time
	Quantity     float64
	BoughtAt     float64
	Provider     string
}

type Portfolio struct {
	Holdings []Holding
}

type LSEPrice struct {
	Close float64 `json:"lastclose"`
}
