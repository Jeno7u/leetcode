package main


// я думаю, что решение через numCount hash map это лучшее
// а ниже алгоритм бойера - мура
func majorityElement(nums []int) int {
    res := nums[0]
    count := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] != res {
            count--
        } else {
            count++
        }

        if count == 0 {
            res = nums[i]
            count++
        }
    }
    return res
}