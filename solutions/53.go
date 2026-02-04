package main

// если меньше нуля то сбрасываем счетчик. НО нужно учитывать кейс когда все числа минимальные и тогда надо найти максимальное отрицательное число
// это как раз учитывается в 14 строчке
func maxSubArray(nums []int) int {
    result := nums[0]
    curr := 0
    for i := range nums {
        if curr < 0 {
            curr = 0
        }

        curr += nums[i]
        result = max(result, curr)
    }
    return result
}