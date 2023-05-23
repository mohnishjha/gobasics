//A small Banking program using methods in Go

package main

import "fmt"

type Account struct {
	accountNumber string
	balance       float64
}

func (a *Account) deposit(amount float64) float64 {
	a.balance += amount
	return a.balance
}

func (a *Account) witdraw(amount float64) float64 {
	a.balance -= amount
	return a.balance
}

func (a *Account) chkbalance() float64 {
	return a.balance
}

func main() {

	cus := Account{
		accountNumber: "123",
		balance:       5000,
	}

	cus2 := Account{
		accountNumber: "124",
		balance:       8000,
	}

	fmt.Println("Deposited 2000 in Cus1's account. Current Balance is: ", cus.deposit(2000), "The account no. is: ", cus.accountNumber)
	fmt.Println("Withdrew 1000 in cus1's account. Current Balance is: ", cus.witdraw(1000), "The accounot no. is: ", cus.accountNumber)
	fmt.Println("Current Balance is: ", cus.chkbalance())

	fmt.Println("Deposited 12000 in Cus2's account. Current Balance is: ", cus2.deposit(12000), "The account no. is: ", cus.accountNumber)
	fmt.Println("Withdrew 9500 in cus2's account. Current Balance is: ", cus2.witdraw(9500), "The account no. is: ", cus2.accountNumber)
	fmt.Println("Current Balance for cus2 is: ", cus2.balance)

}
