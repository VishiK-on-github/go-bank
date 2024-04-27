package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	ToAccount int   `json:"toAccount"`
	Amount    int64 `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string
	LastName  string
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstname string, lastname string) *Account {
	return &Account{
		FirstName: firstname,
		LastName:  lastname,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}

func NewTransfer(toAccount int, amount int64) *TransferRequest {
	return &TransferRequest{
		ToAccount: toAccount,
		Amount:    amount,
	}
}
