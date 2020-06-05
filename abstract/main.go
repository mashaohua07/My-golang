package main

import (
	"fmt"
)

type Account struct {
	AccountNo string
	Pass string
	Balance float64
}

func (account *Account) Deposite(money float64, pass string)  {
	if pass != account.Pass {
		fmt.Println("you password error")
		return
	}

	if money < 0 {
		fmt.Println("you Amount error")
		return
	}

	account.Balance += money
	fmt.Println("success")
}

func (account *Account) WithDraw(money float64, pass string)  {
	if pass != account.Pass {
		fmt.Println("you password error")
		return
	}

	if money < 0 || money > account.Balance {
		fmt.Println("you Amount error")
		return
	}

	account.Balance -= money
	fmt.Println("success")
}

func (account *Account) Query(pass string) {
	if pass != account.Pass {
		fmt.Println("you password error")
		return
	}

	fmt.Println("AccountNo:", account.AccountNo, "Balance:", account.Balance)
}

func main()  {
	account := &Account{
		AccountNo: "gs111111",
		Pass: "6666",
		Balance: 100.0,
	}
	account.Query("6666")
}
