package main


func longestSubarray(nums []int) int {
    result := 0
    idx0, l, r := -1, 0, 0

    for r < len(nums) {
        if nums[r] == 0 {
            if idx0 != -1 {
                result = max(result, r - l - 1)
                l = idx0 + 1
            }
            idx0 = r
        }
        r++
    }
    result = max(result, r - l - 1)
    return result
}