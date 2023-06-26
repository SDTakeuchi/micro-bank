package model

import "time"

type AccountTransaction struct {
	ID        int64
	AccountID string
	Amount    float64
	// Currency  Currency
	CreatedAt time.Time
}
