package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	got := wallet.Deposit(20.50)
	want := 20.50

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
