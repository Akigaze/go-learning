package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("test Deposit", func(t *testing.T) {
		wallet := Wallet{}
		// 创建一个自定义类型的变量，使用 类型(实际值)
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("test Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		// 创建一个自定义类型的变量，使用 类型(实际值)
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, &wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("test Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		// 创建一个自定义类型的变量，使用 类型(实际值)
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, &wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})

}

func assertBalance(t *testing.T, wallet *Wallet, wanted Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != wanted {
		// 想要使用String(), 需要使用 %s 进行格式化
		t.Errorf("expect balance is %s, but actually it's %s", wanted, got)
	}
}

func assertNoError(t *testing.T, get error) {
	t.Helper()

	if get != nil {
		t.Fatal("got an error but didn't wanted")
	}
}

func assertError(t *testing.T, get, wanted error) {
	t.Helper()

	if get == nil {
		t.Fatal("wanted an error but didn't got")
	}
	if get != wanted {
		t.Errorf("expect error '%s', but actually got '%s'", wanted, get)
	}
}
