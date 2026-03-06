package main

import "testing"

func TestAccount(t *testing.T) {
	test_account := Account{
		balance: 500,
		account_owner: Customer{
			name:     "Michail",
			lastname: "Bredichin",
			phone:    812384884,
		},
		account_number: "ABNA134124566132",
	}

	t.Run("Test deposit to the bank account", func(t *testing.T) {
		dep := 100.00
		cb := test_account.Balance()
		err := test_account.Deposit(dep)

		if err != nil {
			t.Error("Got error depositing amount, expect none")
		} else {
			if test_account.Balance() != cb+dep {
				t.Error("Balance after deposit does not match")
			}
		}
	})

	t.Run("Test invalid input", func(t *testing.T) {
		err := test_account.Deposit(-100)
		if err == nil {
			t.Error("Expect to handle negative deposit , should not be possible ")
		}
	})

	t.Run("Negative balance", func(t *testing.T) {
		cb := test_account.Balance()
		err := test_account.Withdraw(cb * 2)
		if err == nil {
			t.Error("Should not be possible to withdraw more than balance")
		}
	})
}
