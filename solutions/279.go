package main


// идентична задачи Coin Change Leetcode 322
func numSquares(n int) int {
    squares := []int{}
    for i := 1; i < n + 1; i++ {
        square := i * i
        if square > n {
            break
        }
        squares = append(squares, square)
    }

    dp := make([]int, n + 1)
    for i := range dp {
        dp[i] = n + 1
    }
    dp[0] = 0

    for a := 1; a < n + 1; a++ {
        for _, square := range squares {
            if a - square >= 0 {
                dp[a] = min(dp[a], 1 + dp[a - square])
            }
        }
    }
    return dp[n]
}