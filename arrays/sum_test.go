package arrays

import "testing"

func TestSum(t *testing.T) {
	nums := [5]int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 20

	if got != want {
		t.Errorf("Got %d want %d given %v", got, want, nums)
	}

}
