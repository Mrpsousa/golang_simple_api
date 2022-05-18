package entity

import (
	"errors"
)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		err := errors.New("you dont have limit for this transaction")
		return err
	}
	if t.Amount < 1 {
		err := errors.New("the amount must be greater than 1")
		return err
	}
	return nil

}
