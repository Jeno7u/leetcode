package main

import "slices"

// самое очевидное ршеение
func topKFrequent1(nums []int, k int) []int {
    // получить частоты list = [(val, freq)] 
    freq := map[int]int{}
    for i := range nums {
        freq[nums[i]]++
    }

    freqSlice := make([][2]int, 0, len(freq))
    for val, freq := range freq {
        freqSlice = append(freqSlice, [2]int{val, freq})
    }

    // отсортировать по частотам и взять первые k
    slices.SortFunc(freqSlice, func(a, b [2]int) int {
        return b[1] - a[1]
    })

    result := make([]int, k)
    for i := 0; i < k; i++ {
        result[i] = freqSlice[i][0]
    }
    return result
}


// а здесь мы создаем freqSlice в который мы сохраняем элементы беря частоту (-1) за индекс
// то есть для каждой частоты у нас список чисел имеющие такую частоту. И потом идем в обратном порядке добавляя числа из списка
// 1, 	        2, 3,    4
// [val1, val2], ,[val3],
func topKFrequent2(nums []int, k int) []int {
    freq := map[int]int{}
    for i := range nums {
        freq[nums[i]]++
    }

    freqSlice := make([][]int, len(nums))

    for val, freq := range freq {
        freqSlice[freq - 1] = append(freqSlice[freq - 1], val)
    }

    result := []int{}
    for i := len(nums) - 1; i >= 0; i-- {
        for _, val := range freqSlice[i] {
            result = append(result, val)
			if len(result) == k {
            	return result
        	}
        }
    }
    return result
}