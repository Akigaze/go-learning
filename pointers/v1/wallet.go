package pointers

import "fmt"

type Wallet struct {
	// balance 小写字母开头的属性 是私有属性
	balance int
}

// Deposit 调用结构体方法时，执行器会将结构体对象传入方法中，这与函数的参数列表传参是一样的道理
// 默认情况下传递的是整个对象的拷贝，如果要修改实参，需要用指针类型作为形参
func (w Wallet) Deposit(amount int) {
	fmt.Println("address of Wallet.balance in Deposit function:", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
