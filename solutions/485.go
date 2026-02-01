package main


func findMaxConsecutiveOnes(nums []int) int {
    result := 0
    l, r := 0, 0
    for r < len(nums) {
        if nums[r] == 0 {
            result = max(result, r - l)
            l = r + 1
        }
        r++
    }
    return max(result, r - l)
}