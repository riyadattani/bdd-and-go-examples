package bank

type Account struct {
	Balance int
}

func (a *Account) GetBalance() int {
	return a.Balance
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) {
	a.Balance -= amount
}

