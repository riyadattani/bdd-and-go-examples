package bank_test

import (
	"bdd-and-go-examples/pkg/bank"
	"testing"
)

func TestBankAccount_Deposit(t *testing.T) {
	t.Run("When I deposit 10 pounds into my account, I expect my balance to increase by 10 pounds", func(t *testing.T) {
		expectedAccountBalance := 10

		account := bank.Account{Balance: 0}

		account.Deposit(expectedAccountBalance)
		accountBalance := account.GetBalance()

		if accountBalance != expectedAccountBalance {
			t.Fatalf("expected %v, got %v", expectedAccountBalance, accountBalance)
		}
	})

	t.Run("When I withdraw 10 pounds from my account, I expect my balance to decrease by 10 pounds", func(t *testing.T) {
		expectedAccountBalance := 5

		account := bank.Account{Balance: 15}

		account.Withdraw(10)
		accountBalance := account.GetBalance()

		if accountBalance != expectedAccountBalance {
			t.Fatalf("expected %v, got %v", expectedAccountBalance, accountBalance)
		}
	})
}
