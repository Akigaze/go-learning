package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	// 创建一个自定义类型的变量，使用 类型(实际值)
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	wanted := Bitcoin(11)

	if got != wanted {
		// 想要使用String(), 需要使用 %s 进行格式化
		t.Errorf("expect balance is %s, but actually it's %s", wanted, got)
	}
}
