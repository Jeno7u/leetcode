package main

// по сути является половиной задачи Leetcode 33. Мы просто делаем lower bound binary search с условием <= nums[len(nums) - 1]
func findMin(nums []int) int {
    l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] <= nums[len(nums) - 1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}
