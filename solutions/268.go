package main

func missingNumber(nums []int) int {
    seen := make([]bool, len(nums) + 1)

    for i := range nums {
        seen[nums[i]] = true
    }

    for i := range seen {
        if seen[i] == false {
            return i
        }
    }
    return -1
}

func missingNumberCustom(nums []int) int {
    n := len(nums) 
    expectedSum := n * (n + 1) / 2
    for i := range nums {
        expectedSum -= nums[i]
    }
    return expectedSum
}