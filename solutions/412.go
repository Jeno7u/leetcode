package main

import "strconv"

func fizzBuzz(n int) []string {
    result := make([]string, n)
    for i := range n {
        if (i + 1) % 3 == 0 && (i + 1) % 5 == 0 {
            result[i] = "FizzBuzz"
        } else if (i + 1) % 3 == 0 {
            result[i] = "Fizz"
        } else if (i + 1) % 5 == 0 {
            result[i] = "Buzz"
        } else {
            result[i] = strconv.Itoa(i + 1)
        }
    }
    return result
}