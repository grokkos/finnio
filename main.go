/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"fmt"
	"github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("API_KEY"))
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

	res, _, _ := finnhubClient.Quote(context.Background()).Symbol("AAPL").Execute()
	//TODO lets move logic and calculations to app and play there
	//TODO begin unit testing
	fmt.Printf("%+v\n", *res.Pc)
}
