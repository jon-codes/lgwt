package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

var ErrInsufficientFunds = errors.New("cannot withdraw due to insufficient funds")

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if w.balance-amt < 0 {
		return ErrInsufficientFunds
	}
	w.balance -= amt
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
