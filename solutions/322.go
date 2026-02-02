package main


// идея в том что мы с 0 до amount считаем минимальное кол-во монет для получения определенной
// amount. Для 0 нужно 0 монет, а в остальном 1 + dp[i - coin] (если i - coin >= 0s)
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
    for i := range dp {
        dp[i] = amount + 1
    }
    dp[0] = 0

    for i := 0; i < amount + 1; i++ {
        for _, coin := range coins {
            if i - coin >= 0 {
                dp[i] = min(dp[i], 1 + dp[i - coin])
            }
        }
    }
    result := dp[amount]
    if result == amount + 1{
        return -1
    }
    return result
}