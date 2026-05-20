// Package pointers Lesson
package pointers

type Wallet struct {
	balance float64
}

func (w Wallet) Deposit(amount float64) float64 {
	return w.balance + amount
}

func (w Wallet) Balance() float64 {
	return w.balance
}
