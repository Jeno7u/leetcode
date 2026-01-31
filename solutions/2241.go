package main


type ATM struct {
    money [5]int
    cash [5]int
}


func YetAnotherConstructor() ATM {
    return ATM{money: [5]int{20, 50, 100, 200, 500}}
}


func (this *ATM) Deposit(banknotesCount []int)  {
    for i := 0; i < 5; i++ {
        this.cash[i] += banknotesCount[i]
    }
}


func (this *ATM) Withdraw(amount int) []int {
    tmp := this.cash
    result := make([]int, 5)
    for i := 4; i >= 0; i-- {
        cashCount := amount / this.money[i] 
        cashCount = min(cashCount, this.cash[i])
        result[i] = cashCount
        this.cash[i] -= cashCount
        amount -= this.money[i] * cashCount
        if amount == 0 {
            return result
        } else if amount < 0 {
            this.cash = tmp
            return []int{-1}
        }
    }
    this.cash = tmp
    return []int{-1}
}
