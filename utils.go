package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
)

func readPortfolioFromFile(filename string) Portfolio {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var portfolio Portfolio
	err = json.Unmarshal(content, &portfolio)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return portfolio
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
