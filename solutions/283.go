package main

// проходимся двумя поинтерами. j указывает на место где все числа до него не нулевые, то есть это место вставки
func moveZeroes(nums []int)  {
    j := 0 // индекс куда вставлять значение не нулевое
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++ 
		}
	}
}
