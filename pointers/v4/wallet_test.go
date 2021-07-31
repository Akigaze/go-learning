package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet *Wallet, wanted Bitcoin) {
		got := wallet.Balance()
		if got != wanted {
			// 想要使用String(), 需要使用 %s 进行格式化
			t.Errorf("expect balance is %s, but actually it's %s", wanted, got)
		}
	}

	t.Run("test Deposit", func(t *testing.T) {
		wallet := Wallet{}
		// 创建一个自定义类型的变量，使用 类型(实际值)
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("test Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		// 创建一个自定义类型的变量，使用 类型(实际值)
		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, &wallet, Bitcoin(10))
	})

}
