package playground

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}

		wallet.Withdraw(Bitcoin(20))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
