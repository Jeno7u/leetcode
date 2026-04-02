package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	maxSum := math.MinInt32
	curr := 0
	for i := 0; i < k; i++ {
		curr += nums[i]
	}
	maxSum = max(maxSum, curr)
	for i := 0; i < len(nums)-k; i++ {
		curr -= nums[i]
		curr += nums[i+k]
		maxSum = max(maxSum, curr)
	}

	return float64(maxSum) / float64(k)

}
