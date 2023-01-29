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
	cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("API_KEY"))
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

	appleData, _, _ := finnhubClient.Quote(context.Background()).Symbol("AAPL").Execute()

	entityApple := &entities.Price{Current: *appleData.C, Previous: *appleData.Pc}

	//use the calcProfitAndLoss function to do the maths
	entityApple.Loss, entityApple.Profit = calcProfitAndLoss(*appleData.C, *appleData.Pc)
	fmt.Printf("%+v\n", entityApple)

	MicrosoftData, _, _ := finnhubClient.Quote(context.Background()).Symbol("MSFT").Execute()
	entityMicrosoft := &entities.Price{Current: *MicrosoftData.C, Previous: *MicrosoftData.Pc}

	//use the calcProfitAndLoss function to do the maths
	entityMicrosoft.Loss, entityMicrosoft.Profit = calcProfitAndLoss(*MicrosoftData.C, *MicrosoftData.Pc)
	fmt.Printf("%+v\n", entityMicrosoft)
}

func calcProfitAndLoss(current float32, previousClosing float32) (loss float32, profit float32) {
	if current > previousClosing {
		profit = current - previousClosing
	} else if current < previousClosing {
		loss = previousClosing - current
	}
	return
}
