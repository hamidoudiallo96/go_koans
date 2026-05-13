package integers

import(
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	got := Add(81, 729)
	fmt.Println(got)
	want := 810
	// Output: 6

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
