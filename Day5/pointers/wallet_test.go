package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, w Wallet, expectedBalance Bitcoin){
		t.Helper()

		balance := w.Balance()
		if balance != expectedBalance {
			t.Errorf("Expected %v but was %v", expectedBalance, balance)
		}	
	}

	assertError := func(t *testing.T, expected, actual error){
		t.Helper()

		if expected == nil {
			t.Fatal("Expected error but was nil")	
		}

		if expected != actual {
			t.Errorf("Error message expected to contain %q but was %q", expected, actual)
		}
	}

	assertNoError := func(t *testing.T, err error){
		t.Helper()

		if err != nil {
			t.Errorf("Error not expected error but was %q", err)	
		}
	}

	t.Run("Deposit_Bitcoin_ReturnsBalance", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
	
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw_EnoughBalance_ReturnsBalance", func(t *testing.T){
		wallet := Wallet{balance: 20}

		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw_NotEnoughBalance_ReturnsErrors", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}

		error := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, error, ErrInsufficientFunds)
	})
}
