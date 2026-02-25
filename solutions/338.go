package main

// мой вариант завязан на том что при увеличении кол-ва бит в числе на 1, число выглядят также
// как и раньше, но с добавлением 1 спереди. Если хранить текущую степень (+ 1) можно получить
// кол-во бит в числе без 1 спереди. dynamic programming.

// есть так же вариант, где если число четное, то берем кол-во бит result[i / 2]. 
// Если нечетное, то берем кол-во бит в предыдущем + 1
func countBits(n int) []int {
    if n == 0 {
        return []int{0}
    }
    if n == 1 {
        return []int{0, 1}
    }
    
    result := make([]int, n + 1)
    result[0] = 0
    result[1] = 1

    bitLength := 2
    remain := bitLength
    for i := 2; i < n + 1; i++ {
        result[i] = 1 + result[i - bitLength]
        remain--
        if remain == 0 {
            bitLength *= 2
            remain = bitLength
        }
    }
    return result
}