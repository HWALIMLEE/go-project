package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account { // Account를 살펴보고 있음(point to Account)
	account := Account{owner: owner, balance: 0}
	return &account
}

// Depost x amount on your account
func (a *Account) Deposit(amount int) { //receiver
	a.balance += amount
}

// Balance of your account
func (a *Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
// error 를 체크하도록 강제시킴
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string { // go 가 자동으로 호출시키는 method
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
