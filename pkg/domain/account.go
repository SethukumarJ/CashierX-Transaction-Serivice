package domain

import (
	"time"
)

type AccountType string

const (
	Savings      AccountType = "savings"
	Current      AccountType = "current"
	FixedDeposit AccountType = "fixed deposit"
)

type Accounts struct {
	AccountID     uint        `json:"account_id" gorm:"unique;autoIncrement"`
	UserID        int64       `json:"user_id" gorm:"column:user_id"`
	AccountNumber int64       `json:"account_number" gorm:"unique"`
	AccountHolder string      `json:"account_holder" gorm:"column:account_holder;type:varchar(100)"`
	AccountType   AccountType `json:"account_type" gorm:"column:account_type"`
	Balance       float64     `json:"balance" gorm:"column:balance"`
	CreatedAt     time.Time   `json:"created_at" gorm:"column:created_at"`
}

// type Transactions struct {
// 	TransactionID   uint      `json:"transaction_id" gorm:"column:transaction_id;unique;autoIncrement"`
// 	AccountID       uint      `json:"account_id" gorm:"not null"`
// 	TransactionType string    `json:"transaction_type" gorm:"not null;type:ENUM('deposit', 'withdrawal', 'transfer')"`
// 	Amount          float64   `json:"amount" gorm:"not null"`
// 	BalanceBefore   float64   `json:"balance_before" gorm:"not null"`
// 	BalanceAfter    float64   `json:"balance_after" gorm:"not null"`
// 	TransferredFrom *uint     `json:"transferred_from,omitempty"`
// 	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
// }
