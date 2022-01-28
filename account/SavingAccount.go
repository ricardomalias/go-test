package account

import customer "go-test/customer"

type SavingAccount struct {
	Owner         customer.Owner
	AgencyNumber  int
	AccountNumber int
	Operation     int
	balance       float64
}

func (c *SavingAccount) Withdraw(amount float64) (string, float64) {
	if amount <= c.balance && amount > 0 {
		c.balance -= amount
		return "Withdraw done", c.balance
	}

	return "Not enough balance", c.balance
}

func (c *SavingAccount) Deposit(amount float64) (string, float64) {
	if amount > 0 {
		c.balance += amount
		return "Deposit done", c.balance
	}

	return "Amount invalid", c.balance
}

func (account *SavingAccount) GetBalance() float64 {
	return account.balance
}
