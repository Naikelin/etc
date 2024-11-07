package models

// Price represents a stock price at a specific date.
type Price struct {
	Date  string  `json:"date"`
	Price float64 `json:"price"`
}
