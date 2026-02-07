package main


func canJump(nums []int) bool {
    jumpLength := nums[0]
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] > jumpLength {
            jumpLength = nums[i]
        }
        jumpLength--
        if jumpLength < 0 {
            return false
        }
    }
    return true
}