// Package array handles array functionality
package array

func Sum(nums []int) int {
	total := 0

	for _, number := range nums {
		total += number
	}

	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var tailSums []int
	for _, number := range numbersToSum {
		if len(number) == 0 {
			tailSums = append(tailSums, 0)
			continue
		}
		tail := number[1:]
		tailSums = append(tailSums, Sum(tail))
	}

	return tailSums
}
