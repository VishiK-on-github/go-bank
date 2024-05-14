package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Token  string `json:"token"`
	Number int64  `json:"number"`
}

type LoginRequest struct {
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type TransferRequest struct {
	ToAccount int   `json:"toAccount"`
	Amount    int64 `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstname"`
	LastName          string    `json:"lastname"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (a *Account) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

func NewAccount(firstname string, lastname string, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:         firstname,
		LastName:          lastname,
		EncryptedPassword: string(encpw),
		Number:            int64(rand.Intn(1000000)),
		CreatedAt:         time.Now().UTC(),
	}, nil
}

func NewTransfer(toAccount int, amount int64) *TransferRequest {
	return &TransferRequest{
		ToAccount: toAccount,
		Amount:    amount,
	}
}
