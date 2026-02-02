package main

// надо просто аккуратно с границами
func reverse(x int) int {
    isNegative := false
    if x < 0 {
        isNegative = true
        x = -x
    }
    reversed := 0
    for x > 0 {
        digit := x % 10
        x /= 10
        if (reversed > (2147483647) / 10) || (reversed == (2147483647) / 10 && digit > 7){
            return 0
        }
        reversed = (reversed * 10) + digit
    }
    if isNegative {
        return -reversed
    } 
    return reversed
}