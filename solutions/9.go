package main


func isPalindromeNumbers(x int) bool {
    if x < 0 {
        return false
    }

    reverse, xcopy := 0, x

    for x > 0 {
        reverse = (reverse * 10) + (x % 10)
        x /= 10
    }
    return reverse == xcopy
}