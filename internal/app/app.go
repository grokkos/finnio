package app

import (
	"context"
	"fmt"
	"github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/grokkos/finnio/internal/pkg/constants"
	"github.com/grokkos/finnio/internal/pkg/entities"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

func App() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		//fmt.Println(".env file loaded with success")
	}
	cfg := finnhub.NewConfiguration()
	//add the api key in header
	cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("API_KEY"))
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

	//TODO refactor to use goroutines for the outgoing requests to finnhub api
	//perform the authenticated api request to retrieve data
	appleData, _, _ := finnhubClient.Quote(context.Background()).Symbol(string(constants.ShareSymbolApple)).Execute()
	//use the entity Price to store all the relevant data we need
	entityApple := entities.Price{Current: *appleData.C, PreviousClosing: *appleData.Pc, Portfolio: *appleData.C * 10}

	//perform the authenticated api request to retrieve data
	microsoftData, _, _ := finnhubClient.Quote(context.Background()).Symbol(string(constants.ShareSymbolMicrosoft)).Execute()
	//use the entity Price to store all the relevant data we need
	entityMicrosoft := entities.Price{Current: *microsoftData.C, PreviousClosing: *microsoftData.Pc, Portfolio: *microsoftData.C * 10}

	//using two goroutines for the calculations and print to console
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		entityApple.Loss, entityApple.Profit = CalcProfitAndLoss(*appleData.C, *appleData.Pc)
		fmt.Printf("AAPL: %+v\n", entityApple)
		wg.Done()
	}()
	go func() {
		entityMicrosoft.Loss, entityMicrosoft.Profit = CalcProfitAndLoss(*microsoftData.C, *microsoftData.Pc)
		fmt.Printf("MSFT: %+v\n", entityMicrosoft)
		wg.Done()
	}()
	wg.Wait()

}

func CalcProfitAndLoss(current float32, previousClosing float32) (loss float32, profit float32) {
	if current > previousClosing {
		profit = (current - previousClosing) * 10
	} else if current < previousClosing {
		loss = (previousClosing - current) * 10
	}
	return
}
