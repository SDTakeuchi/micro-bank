package model

import "time"

type Transfer struct {
	ID                   string
	Currency             Currency
	Amount               float64
	SenderAccountID      string
	ReceiverAccountID    string
	RequestedAt          time.Time
	CompletedAt          *time.Time
}
