package main

import (
	"fmt"

	"github.com/hwalim/go-project/bank_and_dictionary_proj/banking/banking"
)

func main() {
	account := banking.Account{Owner: "hwalim", Balance: 1000}
	fmt.Println(account)
}
