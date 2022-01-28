package main

import (
	"fmt"
	Account "go-test/account"
	Customer "go-test/customer"
)

func main() {
	customer1 := Customer.Owner{Name: "Geroudo", Document: "123.456.789-00", Occupation: "Painter"}
	customer2 := Customer.Owner{Name: "Abigobaldo", Document: "231.456.789-00", Occupation: "Mason"}

	account1 := Account.CheckingAccount{Owner: customer1, AgencyNumber: 589, AccountNumber: 123456}
	account2 := Account.CheckingAccount{Owner: customer2, AgencyNumber: 543, AccountNumber: 46543}

	account1.Deposit(125.5)
	account2.Deposit(275.3)
	fmt.Println("account1: ", account1.GetBalance())
	fmt.Println("account2: ", account2.GetBalance())

	account1.Withdraw(100)
	account1.Deposit(400)
	fmt.Println("account1: ", account1.GetBalance())
	fmt.Println("account2: ", account2.GetBalance())

	fmt.Println(account2.GetBalance())
	fmt.Println(account1.Transfer(150, &account2))
	fmt.Println("account1: ", account1.GetBalance())
	fmt.Println("account2: ", account2.GetBalance())
}
