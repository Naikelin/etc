package models

import (
	"fmt"
	"math"
	"time"
)

// Portfolio represents a collection of stocks in the portfolio.
type Portfolio struct {
	Stocks []Stock `json:"portfolio"`
}

// Profit calculates the profit between two dates for a portfolio.
// where:
// Profit = FinalValue - InitialValue
func (p *Portfolio) Profit(startDate, endDate time.Time) (float64, error) {
	var initialValue, finalValue float64

	// Iterate over all stocks in the portfolio
	for _, stock := range p.Stocks {
		startPrice, err := stock.getPrice(startDate)
		if err != nil {
			return 0, err
		}

		endPrice, err := stock.getPrice(endDate)
		if err != nil {
			return 0, err
		}

		// Check for negative prices
		if startPrice < 0 || endPrice < 0 {
			return 0, fmt.Errorf("price cannot be negative, start price: %.2f, end price: %.2f", startPrice, endPrice)
		}

		// Calculate total value at the start and at the end
		initialValue += float64(stock.Quantity) * startPrice
		finalValue += float64(stock.Quantity) * endPrice
	}

	// Calculate profit
	profit := finalValue - initialValue
	return profit, nil
}

// AnnualizedReturn calculates the annualized return between two dates for a portfolio.
// where:
// AnnualizedReturn = (FinalValue / InitialValue)^(1 / Years) - 1
// Years = (EndDate - StartDate)[days] / 365
func (p *Portfolio) AnnualizedReturn(startDate, endDate time.Time) (float64, error) {
	var initialValue, finalValue float64

	// Iterate over all stocks in the portfolio
	for _, stock := range p.Stocks {
		startPrice, err := stock.getPrice(startDate)
		if err != nil {
			return 0, err
		}

		endPrice, err := stock.getPrice(endDate)
		if err != nil {
			return 0, err
		}

		if startPrice < 0 || endPrice < 0 {
			return 0, fmt.Errorf("price cannot be negative, start price: %.2f, end price: %.2f", startPrice, endPrice)
		}

		initialValue += float64(stock.Quantity) * startPrice
		finalValue += float64(stock.Quantity) * endPrice
	}

	if initialValue == 0 {
		return 0, fmt.Errorf("initial value is zero, cannot calculate return")
	}

	years := (endDate.Sub(startDate).Hours() / 24) / 365

	fmt.Println(years)

	if years <= 0 {
		return 0, fmt.Errorf("invalid time period, cannot calculate annualized return")
	}

	annualizedReturn := math.Pow(finalValue/initialValue, 1/years) - 1

	return annualizedReturn, nil
}
