package fintech

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin // lowercase means private attribute hahaha yessssss
}

type Stringer interface {
	String() string
}

var ErrorInsufficientFunds = errors.New("insufficient funds") // var defines a global variable for the package

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrorInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() string {
	return w.balance.String()
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}