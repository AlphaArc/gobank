package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct{
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UpdateAccountRequest struct{
	ID        int 		`json:"id"`
	FirstName string 	`json:"firstName"`
	LastName  string 	`json:"lastName"`
	Number    int64 	`json:"number"`
	Balance   int64 	`json:"balance"`
}

type Account struct {
	ID        int 		`json:"id"`
	FirstName string 	`json:"firstName"`
	LastName  string 	`json:"lastName"`
	Number    int64 	`json:"number"`
	Balance   int64 	`json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID: 		rand.Intn(10000),
		FirstName: 	firstName,
		LastName: 	lastName,
		Number: 	int64(rand.Intn(10000000)),
		CreatedAt: 	time.Now().UTC(),
	}
}