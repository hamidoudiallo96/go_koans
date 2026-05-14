package arrays

import "fmt"

func Sum(nums [5]int) int {
	var total int

	for i := range len(nums) {
		total += nums[i]
	}

	fmt.Println("Total: ", total)
	return total
}
