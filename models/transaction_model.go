package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Id          int
	Txn_date    time.Time
	Who         string
	Description string
	Payee       string
	Amount      decimal.Decimal
	Category    string
}
