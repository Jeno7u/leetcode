package main

// создаем prefix произведения и suffix. Любое число в answer можно представить как значение префикса до значения *
// на значение suffix после значения.
func productExceptSelfExtraMemory(nums []int) []int {
    prefixProduct := make([]int, len(nums) - 1)
    suffixProduct := make([]int, len(nums) - 1)

    prefixProduct[0] = nums[0]
    suffixProduct[len(nums) - 2] = nums[len(nums) - 1]
    for i := 1; i < len(nums) - 1; i++ {
        prefixProduct[i] = prefixProduct[i - 1] * nums[i]
    }
    for i := len(nums) - 2; i >= 1; i-- {
        suffixProduct[i - 1] = suffixProduct[i] * nums[i]
    }

    answer := make([]int, len(nums))
    answer[0] = suffixProduct[0]
    answer[len(nums) - 1] = prefixProduct[len(nums) - 2]
    for i := 1; i < len(nums) - 1; i++ {
        answer[i] = prefixProduct[i - 1] * suffixProduct[i]
    }
    
    return answer
}


// та же идея, что и выше, но мы храним префиксные и суффиксные суммы внутри answer и перемножаем
func productExceptSelf(nums []int) []int {
    answer := make([]int, len(nums))
    answer[0] = 1

    curr := nums[0]
    for i := 0; i < len(nums) - 1; i++ {
        answer[i + 1] = curr
        curr *= nums[i + 1]
    }
    curr = nums[len(nums) - 1]
    for i := len(nums) - 2; i >= 0; i-- {
        answer[i] *= curr
        curr *= nums[i]
    }
    
    
    return answer
}