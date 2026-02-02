package main

// мы считаем сумму через вычесление [:i] - prefixSum[j]. j мы находим из
// расчета, что если у нас curr = curr % k, то чтобы curr == 0, нам нужно
// вычесть префиксную сумму, где она равна curr. А это мы как раз сохраняем.
// prefixSum[curr] = индекс первого появления (потому что последнее может уменьшить
// размер длины на 1, что в случае если у нас длина текущая 2, то ошибка) +
// учесть, что curr == 0 не может быть на i = 0
func checkSubarraySum(nums []int, k int) bool {
    if len(nums) < 2 {
        return false
    }

    curr := 0
    prefixSum := map[int]int{}
    for i := 0; i < len(nums); i++ {
        curr = (curr + nums[i]) % k
        if curr == 0 && i != 0{
            return true
        } 
        val, ok := prefixSum[curr]
        if ok && val != i - 1{
            return true
        }
        if !ok {
            prefixSum[curr] = i
        }
    }
    return false
}