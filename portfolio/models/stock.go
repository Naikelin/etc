package models

import (
	"fmt"
	"time"
)

// Stock represents a stock with its name, quantity, and prices.
type Stock struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Prices   []Price `json:"prices"`
}

// getPrice returns the price of a stock at a specific date.
func (s *Stock) getPrice(date time.Time) (float64, error) {
	for _, price := range s.Prices {
		priceDate, err := time.Parse("2006-01-02T00:00:00Z", price.Date)
		if err != nil {
			return 0, fmt.Errorf("error parsing date: %v", err)
		}
		if priceDate.Equal(date) {
			return price.Price, nil
		}
	}
	return 0, fmt.Errorf("no price found for date: %s", date)
}
