package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	fmt.Println("address of Wallet.balance in Test function:", &wallet.balance)
	wallet.Deposit(10)
	got := wallet.Balance()
	wanted := 10

	if got != wanted {
		t.Errorf("expect balance is %d, but actually it's %d", wanted, got)
	}
}
