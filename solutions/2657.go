package main

import "fmt"

// вместо hash map/set для определения были ли появления цифр до этого можно исопльзовать список и сохранять по индексу ( число - 1 )
func findThePrefixCommonArray(A []int, B []int) []int {
	result := make([]int, len(A))
	amountCommon := 0
	commonNumbers := make([]int, len(A) + 1) 
	for i := 0; i < len(A); i++ {
        commonNumbers[A[i]]++
		if commonNumbers[A[i]] == 2 {
			amountCommon++
		}

        commonNumbers[B[i]]++
		if commonNumbers[B[i]] == 2 {
			amountCommon++
		}

		result[i] = amountCommon
	}
	
	return result
}

func main() {
	A := []int{1,2,3}
	B := []int{1,3,2}
	fmt.Println(findThePrefixCommonArray(A, B))
}