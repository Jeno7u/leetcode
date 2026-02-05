package main


func maxProduct(nums []int) int {
    result := getMax(nums)
    currMin, currMax := 1, 1
    for i := range nums {
        if nums[i] == 0 {
            currMin, currMax = 1, 1
        }
        tmp := currMax * nums[i]
        currMax = max(tmp, currMin * nums[i], nums[i])
        currMin = min(tmp, currMin * nums[i], nums[i])
        result = max(result, currMax)
    }
    return result
}

func getMax(nums []int) int {
    max := -11
    for i := range nums {
        if max < nums[i] {
            return max
        }
    }
    return max
}