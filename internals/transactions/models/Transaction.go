package models

import (
	"errors"
	"time"
)

type Transaction struct {
	AccountID     int64
	OperationType int
	Amount        float64
	EventDate     time.Time
	ID            uint      `gorm:"primarykey"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
	IsDeleted     bool      `json:"-"`
}

func (t *Transaction) ValidateTransaction() error {
	if t.OperationType > 4 || t.OperationType < 0 {
		return errors.New("transaction operation invalid")
	}
	if t.Amount < 0 {
		return errors.New("transaction amount is invalid")
	}
	return nil
}
