package pointers

import "fmt"

// Bitcoin 自定义类型，type 关键字允许将一种类型定义成其他类型，同时还可以给定义的类型添加方法
type Bitcoin int

// String 方法是在使用fmt打印时，使用 %s 进行格式化时会被调用
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// balance 小写字母开头的属性 是私有属性
	balance Bitcoin
}

// Deposit 使用指针类型作为调用者形参，就相当于接收实参对象，指针类型的形参和非指针类型的形参在使用上是相同的
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
