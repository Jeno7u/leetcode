package main

// добраться до следующей ступеньки можно либо с предыдущей
// либо пред предыдущей. Тогда просто сложим кол-во способов.
// Там кста получается числа в порядке чисел фибоначи
func climbStairs(n int) int {
    if n <= 3 {
        return n
    }

    idx := 4
    val1, val2 := 3, 5
    for idx != n {
        val1, val2 = val2, val1+val2
        idx++
    }
    return val2
}