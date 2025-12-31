package main

import "fmt"


func longestMonotonicSubarray(nums []int) int {
	result := 0
	curr := 1
	increasing := 0
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] < nums[i + 1] {
			if increasing == 1 {
				curr++
			} else {
				increasing = 1
				curr = 2
			}
		} else if nums[i] > nums[i + 1] {
			if increasing == -1 {
				curr++
			} else {
				increasing = -1
				curr = 2
			}
		} else {
			increasing = 0
			curr = 1
		}
		result = max(result, curr)
	}
	return max(result, curr)
}

func main() {
	nums := []int{3,2,1}
	fmt.Println(longestMonotonicSubarray(nums))
}