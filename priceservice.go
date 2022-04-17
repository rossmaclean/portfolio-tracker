package main

import (
	"encoding/json"
	"fmt"
	"github.com/svarlamov/goyhfin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// GetCloseForDate Returns the close price in pence for the specified date, or an earlier date if date is not trading day
var GetCloseForDate = func(ticker string, date time.Time, provider string) float64 {
	switch provider {
	case "LSE":
		return getLsePrice(ticker)
	case "YAHOO":
		hp := getHistoricalPricesForTicker(ticker)
		xhp := getHistoricalPricesForTicker("USDGBP=X")
		exchangeRate := getCloseForDateYahoo(xhp, date).Close
		return getCloseForDateYahoo(hp, date).Close * exchangeRate * 100
	}
	log.Printf("Unable to get %s price for %s on %s", provider, ticker, date)
	return 0
}

var getCloseForDateYahoo = func(quotes []goyhfin.Quote, date time.Time) goyhfin.Quote {
	earliestDate := quotes[0].ClosesAt
	for i := date; i.Equal(earliestDate) || i.After(earliestDate); i = i.Add(-time.Hour * 24) {
		quote := getQuoteOnDate(quotes, i)
		if quote.Close != 0 {
			return quote
		}
	}
	return goyhfin.Quote{}
}

func getQuoteOnDate(quotes []goyhfin.Quote, date time.Time) goyhfin.Quote {
	for _, quote := range quotes {
		qYear, qMonth, qDay := quote.ClosesAt.Date()
		year, month, day := date.Date()
		if qYear == year && qMonth == month && qDay == day {
			return quote
		}
	}
	return goyhfin.Quote{}
}

var getHistoricalPricesForTicker = func(ticker string) []goyhfin.Quote {
	resp, err := goyhfin.GetTickerData(ticker, goyhfin.TenYear, goyhfin.OneDay, false)
	if err != nil {
		fmt.Println("Error fetching Yahoo Finance data:", err)
		panic(err)
	}
	return resp.Quotes
}

var getLsePrice = func(ticker string) float64 {
	url := "https://api.londonstockexchange.com/api/gw/lse/instruments/alldata/" + ticker

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var lsePrice LSEPrice
	json.Unmarshal(bodyBytes, &lsePrice)
	return lsePrice.Close
}
