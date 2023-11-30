package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shopspring/decimal"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Usage: "Get crypto investment recommendations!",
		Commands: []*cli.Command{
			{
				Name:      "invest",
				Usage:     "Get the recommended BTC and ETH investment (70/30 split) for given USD amount.",
				ArgsUsage: "USD",
				Action:    getRecommendedInvestment,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getRecommendedInvestment(cCtx *cli.Context) error {
	inputFundsStr := cCtx.Args().Get(0)
	if inputFundsStr == "" {
		fmt.Println("USAGE:\n\tcrypto-coach invest [command options] USD")
		return nil
	}

	inputUsd, err := decimal.NewFromString(inputFundsStr)
	if err != nil {
		fmt.Printf("Couldn't convert '%s' to decimal.\n", inputFundsStr)
		return err
	}

	cryptoExchangeRates, err := getCryptoExchangeRates()
	if err != nil {
		fmt.Println("Couldn't get crypto exchange rates")
		return err
	}

	recommendation := getInvestmentRecommendation(inputUsd, defaultInvestmentStrategy, cryptoExchangeRates)
	err = printInvestmentRecommendation(recommendation)
	if err != nil {
		fmt.Println("Couldn't print the investment recommendation")
		return err
	}
	return nil
}
