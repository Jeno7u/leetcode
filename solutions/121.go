package main


func maxProfit(prices []int) int {
    buy := prices[0]
    profit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] < buy {
            buy = prices[i]
        } else if prices[i] - buy > profit { // можно без else, но дольше. Как я понимаю, если buy убывает, то профит не будет самым высоким 
            profit = prices[i] - buy
        }
    }
    return profit
}