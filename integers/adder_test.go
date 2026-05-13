package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(81, 729)
	want := 810

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
