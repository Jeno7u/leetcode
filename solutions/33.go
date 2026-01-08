package main

import "fmt"

// находим индекс этого поворота за O(log n)
func searchBound(nums *[]int) int {
	l, r := 0, len(*nums) - 1
	boundary_idx := -1
	for l <= r {
		mid := (l + r) / 2
		if (*nums)[mid] <= (*nums)[len(*nums) - 1] {
			boundary_idx = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return boundary_idx
}

// делаем обычный binary search, но для этого список должен быть отсортированным. И так 
// как наш список отсортированный, но просто со сдвигом, то мы вернем этот сдвиг обратно
// не через создание нового корректного списка, а будем изменять mid idx (который мы считаем
// как для отсортированного) в convertedIndex, который корректен для nums
func search(nums []int, target int) int {
	boundary_idx := searchBound(&nums)

	l, r := 0, len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		convertedIndex := (mid + boundary_idx) % len(nums)
		if nums[convertedIndex] == target {
			return convertedIndex
		}

		if nums[convertedIndex] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{4,5,6,7,0,1,2}
	target := 0
	fmt.Println(search(nums, target))
}