package app

import (
	"context"
	"fmt"
	"github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/grokkos/finnio/internal/pkg/entities"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	//perform the authenticated api request to retrieve data
	appleData, _, _ := finnhubClient.Quote(context.Background()).Symbol("AAPL").Execute()
	//use the entity Price to store all the relevant data we need
	entityApple := entities.Price{Current: *appleData.C, PreviousClosing: *appleData.Pc, Portfolio: *appleData.C * 10}
	//use the calcProfitAndLoss function to do the maths
	entityApple.Loss, entityApple.Profit = calcProfitAndLoss(*appleData.C, *appleData.Pc)

	//perform the authenticated api request to retrieve data
	microsoftData, _, _ := finnhubClient.Quote(context.Background()).Symbol("MSFT").Execute()
	//use the entity Price to store all the relevant data we need
	entityMicrosoft := entities.Price{Current: *microsoftData.C, PreviousClosing: *microsoftData.Pc, Portfolio: *microsoftData.C * 10}
	//use the calcProfitAndLoss function to do the maths
	entityMicrosoft.Loss, entityMicrosoft.Profit = calcProfitAndLoss(*microsoftData.C, *microsoftData.Pc)

	fmt.Printf("AAPL: %+v\n \nMSFT: %+v\n", entityApple, entityMicrosoft)
}

func calcProfitAndLoss(current float32, previousClosing float32) (loss float32, profit float32) {
	if current > previousClosing {
		profit = (current - previousClosing) * 10
	} else if current < previousClosing {
		loss = (previousClosing - current) * 10
	}
	return
}
