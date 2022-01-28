package account

import customer "go-test/customer"

type CheckingAccount struct {
	Owner         customer.Owner
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (c *CheckingAccount) Withdraw(amount float64) (string, float64) {
	if amount <= c.balance && amount > 0 {
		c.balance -= amount
		return "Withdraw done", c.balance
	}

	return "Not enough balance", c.balance
}

func (c *CheckingAccount) Deposit(amount float64) (string, float64) {
	if amount > 0 {
		c.balance += amount
		return "Deposit done", c.balance
	}

	return "Amount invalid", c.balance
}

func (accountOrigin *CheckingAccount) Transfer(amount float64, accountDestiny *CheckingAccount) bool {
	if amount > accountOrigin.balance && amount > 0 {
		return false
	}

	accountOrigin.balance -= amount
	accountDestiny.Deposit(amount)

	return true
}

func (account *CheckingAccount) GetBalance() float64 {
	return account.balance
}
