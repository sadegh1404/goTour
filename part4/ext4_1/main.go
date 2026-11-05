package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Owner   string
	balance int
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withidraw(amount int) error {
	if amount > a.balance {
		return errors.New("insuficient balance")
	}
	a.balance -= amount
	return nil
}

func (a Account) Balance() int {
	return a.balance
}

func main() {
	my_account := Account{"Sadegh", 100}
	my_account.Deposit(203)
	fmt.Println(my_account.Balance())
	my_account.Withidraw(300)
	fmt.Println(my_account.Balance())
	err := my_account.Withidraw(300)
	fmt.Println(err)

}
