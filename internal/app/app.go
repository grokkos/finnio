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
		fmt.Println(".env file loaded with success")
	}
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("API_KEY"))
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

	res, _, _ := finnhubClient.Quote(context.Background()).Symbol("AAPL").Execute()
	//TODO begin unit testing
	t := &entities.Price{Current: *res.C, Previous: *res.Pc}

	//check for profit and fill the response struct object accordingly
	if *res.C > *res.Pc {
		t.Profit = *res.C - *res.Pc
	} else if *res.C < *res.Pc {
		t.Loss = *res.Pc - *res.C
	}

	fmt.Printf("%+v\n", t)

	//previousClosing := *res.Pc MSFT

	//goroutine to check the MSFT price parallel
	//res1, _, _ := finnhubClient.Quote(context.Background()).Symbol("MSFT").Execute()
}
