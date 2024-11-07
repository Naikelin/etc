package models

import (
	"math"
	"testing"
	"time"
)

// TestPortfolioCalculations validates the profit and annualized return calculations.
func TestPortfolioCalculations(t *testing.T) {
	// Test Case 1: A portfolio with two stocks with prices on two different dates
	portfolio := Portfolio{
		Stocks: []Stock{
			{
				Name:     "TEST",
				Quantity: 10,
				Prices: []Price{
					{Date: "2024-01-01T00:00:00Z", Price: 150.0},
					{Date: "2024-06-01T00:00:00Z", Price: 160.0},
				},
			},
			{
				Name:     "FINTUAL",
				Quantity: 5,
				Prices: []Price{
					{Date: "2024-01-01T00:00:00Z", Price: 2500.0},
					{Date: "2024-06-01T00:00:00Z", Price: 2600.0},
				},
			},
		},
	}

	// Start and end dates for the evaluation
	startDate, _ := time.Parse("2006-01-02T00:00:00Z", "2024-01-01T00:00:00Z")
	endDate, _ := time.Parse("2006-01-02T00:00:00Z", "2024-06-01T00:00:00Z")

	// Test: Profit Calculation
	// Formula: Profit = (Final Price - Initial Price) * Quantity
	// This formula calculates the profit for each stock by subtracting the initial price from the final price
	// and multiplying the result by the quantity of shares held
	// Then we sum the profits for all stocks in the portfolio

	// (160-150)*10 + (2600-2500)*5
	expectedProfit := (160.0-150.0)*10 + (2600.0-2500.0)*5
	profit, err := portfolio.Profit(startDate, endDate)
	if err != nil {
		t.Fatalf("Error calculating profit: %v", err)
	}
	if profit != expectedProfit {
		t.Errorf("Expected profit %.2f, got %.2f", expectedProfit, profit)
	}

	// Test: Annualized Return Calculation
	// Formula: Annualized Return = ((Final Value / Initial Value)^(1 / Years)) - 1
	// Where Final Value = Final Price * Quantity, Initial Value = Initial Price * Quantity
	// Years = (End Date - Start Date)[days] / 365
	// This formula calculates the return per year, accounting for the compounding effect over the given time period.

	// 150 * 10 + 2500 * 5
	initialValue := (150.0 * 10) + (2500.0 * 5)
	// 160 * 10 + 2600 * 5
	finalValue := (160.0 * 10) + (2600.0 * 5)
	years := (endDate.Sub(startDate).Hours() / 24) / 365
	expectedAnnualizedReturn := math.Pow(finalValue/initialValue, 1/years) - 1
	annualizedReturn, err := portfolio.AnnualizedReturn(startDate, endDate)
	if err != nil {
		t.Fatalf("Error calculating annualized return: %v", err)
	}
	// Compare the annualized return with a tolerance of 0.001
	if math.Abs(annualizedReturn-expectedAnnualizedReturn) > 0.001 {
		t.Errorf("Expected annualized return %.3f, got %.3f", expectedAnnualizedReturn, annualizedReturn)
	}

	// Test Case 2: Portfolio with a stock where the price does not change
	portfolio2 := Portfolio{
		Stocks: []Stock{
			{
				Name:     "STATIC",
				Quantity: 10,
				Prices: []Price{
					{Date: "2024-01-01T00:00:00Z", Price: 100.0},
					{Date: "2024-06-01T00:00:00Z", Price: 100.0},
				},
			},
		},
	}

	// No change in price, no profit
	expectedProfit2 := 0.0
	profit2, err := portfolio2.Profit(startDate, endDate)
	if err != nil {
		t.Fatalf("Error calculating profit for STATIC: %v", err)
	}
	if profit2 != expectedProfit2 {
		t.Errorf("Expected profit %.2f for STATIC, got %.2f", expectedProfit2, profit2)
	}

	// Test Case 3: Portfolio with negative or zero prices (Invalid case)
	portfolio3 := Portfolio{
		Stocks: []Stock{
			{
				Name:     "INVALID",
				Quantity: 10,
				Prices: []Price{
					{Date: "2024-01-01T00:00:00Z", Price: -100.0},
					{Date: "2024-06-01T00:00:00Z", Price: 100.0},
				},
			},
		},
	}

	// Expect an error for the negative price
	_, err = portfolio3.Profit(startDate, endDate)
	if err == nil {
		t.Fatalf("Expected error for negative price, but got none")
	}

	// Test Case 4: Portfolio with zero shares (No shares to calculate profit)
	portfolio4 := Portfolio{
		Stocks: []Stock{
			{
				Name:     "NO_SHARES",
				Quantity: 0,
				Prices: []Price{
					{Date: "2024-01-01T00:00:00Z", Price: 100.0},
					{Date: "2024-06-01T00:00:00Z", Price: 100.0},
				},
			},
		},
	}

	// No profit, since there are no shares
	expectedProfit4 := 0.0
	profit4, err := portfolio4.Profit(startDate, endDate)
	if err != nil {
		t.Fatalf("Error calculating profit for NO_SHARES: %v", err)
	}
	if profit4 != expectedProfit4 {
		t.Errorf("Expected profit %.2f for NO_SHARES, got %.2f", expectedProfit4, profit4)
	}

	// Test Case 5: Portfolio with empty data (No stocks, no prices)
	portfolio5 := Portfolio{}
	// No profit to calculate
	expectedProfit5 := 0.0
	profit5, err := portfolio5.Profit(startDate, endDate)
	if err != nil {
		t.Fatalf("Error calculating profit for empty portfolio: %v", err)
	}
	if profit5 != expectedProfit5 {
		t.Errorf("Expected profit %.2f for empty portfolio, got %.2f", expectedProfit5, profit5)
	}
}
