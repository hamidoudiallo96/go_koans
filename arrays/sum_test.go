package array

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15

	if got != want {
		t.Errorf("Got %d instead of %d given %v", got, want, nums)
	}
}

func TestSumAll(t *testing.T) {
	arr1 := []int{1, 2}
	arr2 := []int{0, 9}

	got := SumAll(arr1, arr2)
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("Got %v instead of %v given %v and %v", got, want, arr1, arr2)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		arr1 := []int{16, 88, 97, 92, 34, 98}
		arr2 := []int{12, 5, 77, 7, 11, 13}

		got := SumAllTails(arr1, arr2)
		want := []int{409, 113}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		t.Helper()
		arr1 := []int{}
		arr2 := []int{12, 5, 77, 7, 11, 13}

		got := SumAllTails(arr1, arr2)
		want := []int{0, 113}

		checkSums(t, got, want)
	})
}
