package main

import (
	"errors"
	"fmt"
)

// Implement these methods:
// - `Deposit(amount float64) error`
// - `Withdraw(amount float64) error`
// - `Balance() float64`
// - `String() string` — human-readable summary

// Rules:
// - Negative deposits = error
// - Withdrawing more than balance = error
// - Use pointer receivers where it makes sense — and be intentional about it (you'll be asked why in the review)

type Customer struct {
	name     string
	lastname string
	phone    int
}

type Account struct {
	balance        float64
	account_number string
	account_owner  Customer
}

// we modify balance so pointer receiver makes sense
func (acc *Account) Withdraw(amount float64) error {
	if (acc.balance - amount) >= 0 {
		acc.balance -= amount
		return nil
	}
	return errors.New("negative deposit")
}

// we modify balance so pointer receiver makes sense
func (acc *Account) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("Negative deposite is not possible")
	}
	acc.balance += amount
	return nil
}

// read only so value pointer is fine
func (acc Account) Balance() float64 {
	return acc.balance
}

// read only so value pointer is fine
func (acc Account) Print() string {
	return fmt.Sprintf(" Customer %s %s Balance is : %f", acc.account_owner.name, acc.account_owner.lastname, acc.balance)
}
