package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/shopspring/decimal"
)

type CryptoSymbol string

const coinbaseExchangeRatesEndpoint = "https://api.coinbase.com/v2/exchange-rates?currency=USD"

const (
	Etherium CryptoSymbol = "ETH"
	Bitcoin  CryptoSymbol = "BTC"
)

var cryptoSymbols []CryptoSymbol = []CryptoSymbol{Etherium, Bitcoin}

var defaultInvestmentStrategy InvestmentStrategy = InvestmentStrategy{
	[]InvestmentStrategyProportion{
		{Etherium, 30},
		{Bitcoin, 70},
	},
}

type InvestmentStrategyProportion struct {
	Symbol CryptoSymbol
	Weight int64
}

type InvestmentStrategy struct {
	Proportions []InvestmentStrategyProportion
}

func (strategy InvestmentStrategy) TotalWeight() int64 {
	var totalWeight int64 = 0
	for _, strategyPortion := range strategy.Proportions {
		totalWeight += strategyPortion.Weight
	}
	return totalWeight
}

type InvestmentRecommendation map[CryptoSymbol]decimal.Decimal

type CryptoExchangeRates struct {
	Currency string                           `json:"currency"`
	Rates    map[CryptoSymbol]decimal.Decimal `json:"rates"`
}

type CryptoExchangeRatesResponse struct {
	Data CryptoExchangeRates `json:"data"`
}

func getCryptoExchangeRates() (*CryptoExchangeRates, error) {
	resp, err := http.Get(coinbaseExchangeRatesEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response *CryptoExchangeRatesResponse
	json.Unmarshal(body, &response)
	return &response.Data, nil
}

func getInvestmentRecommendation(fundsUsd decimal.Decimal, strategy InvestmentStrategy, exchangeRates *CryptoExchangeRates) InvestmentRecommendation {
	totalStrategyWeight := decimal.NewFromInt(strategy.TotalWeight())
	recommendation := make(InvestmentRecommendation)

	for _, strategyProportion := range strategy.Proportions {
		symbol := strategyProportion.Symbol
		exchangeRate := exchangeRates.Rates[symbol]
		weightRatio := decimal.NewFromInt(strategyProportion.Weight).Div(totalStrategyWeight)

		recommendedPurchase := fundsUsd.Mul(weightRatio).Mul(exchangeRate)
		recommendation[symbol] = recommendedPurchase
	}
	return recommendation
}

func printInvestmentRecommendation(recommendation InvestmentRecommendation) error {
	jsonData, err := json.Marshal(recommendation)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", jsonData)
	return nil
}
