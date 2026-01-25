package main

// проходимся один раз и на каждом проходе добавляем в словарь текущий результат. Словарь хранит кол-во способов получения определенной суммы из префикса. Когда мы
// проходимся, то мы проверяем что curr - target есть ли в словаре. Ведь сумму любого подсписка можно представить как сумму nums[:i] - префиксная сумма = target
func subarraySum(nums []int, k int) int {
    result := 0
	prefixSum := map[int]int{0: 1}
	curr := 0 
	for i := 0; i < len(nums); i++ {
		curr += nums[i]
		if val, ok := prefixSum[curr - k]; ok {
			result += val
		}

		_, ok := prefixSum[curr]
		if ok {
			prefixSum[curr]++
		} else {
			prefixSum[curr] = 1
		}
	}
	return result
}

