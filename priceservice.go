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

var getCloseForDate = func(quotes []goyhfin.Quote, date time.Time) goyhfin.Quote {
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
	//log.Println(resp.ExchangeName)
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
