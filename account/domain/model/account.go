package model

import "time"

type Account struct {
	ID     string
	UserID string
	Number string
	// BranchID  int64
	Balance   float64
	Currency  Currency
	CreatedAt time.Time
	UpdatedAt time.Time
	BannedAt  time.Time
	DeletedAt time.Time
}
