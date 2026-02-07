package main

// func rob(nums []int) int {
//     dp := make([]int, len(nums))

//     for i := range nums {
//         if i - 2 >= 0 {
//             dp[i] = max(dp[i], nums[i] + dp[i - 2])
//         }
//         if i - 3 >= 0 {
//             dp[i] = max(dp[i], nums[i] + dp[i - 3])
//         }
//         dp[i] = max(dp[i], nums[i])
//     }

//     result := dp[len(dp) - 1]
//     if len(nums) > 1 {
//         result = max(result, dp[len(dp) - 2])
//     }

//     return result
// }

// dp проблема. В целом схожа с Climbing stairs, но нам надо смотреть на максимум между текущий + пред предыдущий
// и текущий + пред пред предыдущий (через один и через два элемента)
func rob(nums []int) int {
    prev1, prev2, prev3, curr := -1, -1, -1, -1
    for i := range nums {
        prev1, prev2, prev3, curr = curr, prev1, prev2, nums[i]
        curr = max(curr, curr + prev2, curr + prev3)
    }

    return max(curr, prev1)
}