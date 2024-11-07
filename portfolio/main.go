package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"portfolio/models"
)

func main() {
	// loads JSON
	data, err := os.ReadFile("./data/portfolio.json")
	if err != nil {
		fmt.Println("Err reading JSON:", err)
		return
	}

	var portfolio models.Portfolio
	err = json.Unmarshal(data, &portfolio)
	if err != nil {
		fmt.Println("Err deserializing JSON:", err)
		return
	}

	startDate, _ := time.Parse("2006-01-02T00:00:00Z", "2024-01-01T00:00:00Z")
	endDate, _ := time.Parse("2006-01-02T00:00:00Z", "2024-06-01T00:00:00Z")

	// calculate profit
	profit, err := portfolio.Profit(startDate, endDate)
	if err != nil {
		fmt.Println("Err calculating profit:", err)
		return
	}
	fmt.Printf("profit between %s and %s: %.2f%%\n", startDate, endDate, profit)

	// calculate annualized return
	annualizedReturn, err := portfolio.AnnualizedReturn(startDate, endDate)
	if err != nil {
		fmt.Println("Err calculating annualized return:", err)
		return
	}
	fmt.Printf("annualized return between %s and %s: %.2f%%\n", startDate, endDate, annualizedReturn*100)
}
