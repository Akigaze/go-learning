package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	wanted := 10

	if got != wanted {
		t.Errorf("expect balance is %d, but actually it's %d", wanted, got)
	}
}
