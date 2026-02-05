package main

// добавляем в стек значения если есть в nums1 и проверяем
// что новое значение из nums2 не больше последнего в стеке.
// если больше, то мы достаем из стека значение и находим куда
// класть текущее
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    nums1Map := map[int]int{}
    for i := range nums1 {
        nums1Map[nums1[i]] = i
    }
    result := make([]int, len(nums1))
    for i := range result {
        result[i] = -1
    }
    stack := []int{}
    for _, curr := range nums2 {
        for len(stack) != 0 && curr > stack[len(stack) - 1] {
            val := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            idx := nums1Map[val]
            result[idx] = curr
        }

        if _, ok := nums1Map[curr]; ok {
            stack = append(stack, curr)
        }
    }
    return result
}