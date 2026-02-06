package main


func minSubArrayLen(target int, nums []int) int {
    result := len(nums) + 1
    curr := 0
    l, r := 0, 0
    for r < len(nums) {
        curr += nums[r]
        
        for curr >= target {
            result = min(result, r - l + 1)
            curr -= nums[l]
            l++
        }
        r++
    }
    if result == len(nums) + 1 {
        return 0
    }
    return result
}