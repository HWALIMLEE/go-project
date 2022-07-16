package main

import (
	"fmt"

	"github.com/hwalim/go-project/bank_and_dictionary_proj/accounts"
)

func main() {
	account := accounts.NewAccount("hwalim")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err) //kill program
		fmt.Println(err)
	}
	fmt.Println(account.Owner())
	account.ChangeOwner("hwalim_new")
	fmt.Println(account.Owner())
	fmt.Println(account)
}
