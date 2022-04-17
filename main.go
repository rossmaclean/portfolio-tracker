package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
)

func main() {
	portfolio := readPortfolioFromFile("holdings.json")
	printMenu(portfolio)
}

func printMenu(portfolio Portfolio) {
	stockReports := createReport(portfolio).StockReports

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "➡ {{ .Ticker | cyan }}",
		Inactive: "  {{ .Ticker | cyan }}",
		Selected: "➡ {{ .Ticker | red | cyan }}",
		Details: `
--------- Holding ----------
{{ "Ticker:" | faint }}	{{ .Ticker }}
{{ "Units Held:" | faint }}	{{ .UnitsHeld }}
{{ "Price (pence):" | faint }}	{{ .Price }}
{{ "Value (£):" | faint }}	{{ .Value }}
{{ "Cost (£):" | faint }}	{{ .Cost }}
-------- Gain/Loss ---------
{{ "£:" | faint }}	{{ .ValueGainLoss }}
{{ "%:" | faint }}	{{ .PercentageGainLoss }}`,
	}

	searcher := func(input string, index int) bool {
		stockReport := stockReports[index]
		name := strings.Replace(strings.ToLower(stockReport.Ticker), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Holdings",
		Items:     stockReports,
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose number %d: %s\n", i+1, stockReports[i].Ticker)
}
