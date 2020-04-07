package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	expected := 10
	balance := wallet.Balance()
	if balance != expected {
		t.Errorf("Expected %d but was %d", expected, balance)
	}
}
