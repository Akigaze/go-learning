package pointers

type Wallet struct {
	// balance 小写字母开头的属性 是私有属性
	balance int
}

// Deposit 使用指针类型作为调用者形参，就相当于接收实参对象，指针类型的形参和非指针类型的形参在使用上是相同的
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
