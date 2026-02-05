package main

// можно легко вычеслить atMostK, но намного сложнее exactly K.
// так что exactlyK = atMostK - atMost(k-1)
func subarraysWithKDistinct(nums []int, k int) int {
    return atMostK(nums, k) - atMostK(nums, k - 1)
}

func atMostK(nums []int, k int) int {
    numCount := map[int]int{}
    l, r := 0, 0
    result := 0
    for r < len(nums) {
        if numCount[nums[r]] == 0 {
            k--
        }
        numCount[nums[r]]++
        for k < 0 {
            numCount[nums[l]]--
            if numCount[nums[l]] == 0 {
                k++
            }
            l++
        }
        result += r - l + 1
        r++
    }
    return result
}