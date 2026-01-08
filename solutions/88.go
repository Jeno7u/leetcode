package main

// по сути это обычный вопрос на merge sorted arrays, но чтоюы решить за O(1) по памяти вместо буфера мы складываем результаты не с начала
// списка как обычно, а заполняем с конца
func merge(nums1 []int, m int, nums2 []int, n int)  {
	p1, p2 := m - 1, n - 1
	idx := len(nums1) - 1
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[idx] = nums1[p1]
			p1--
		} else {
			nums1[idx] = nums2[p2]
			p2--
		}
		idx--
	}
}


func main() {
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3
	merge(nums1, m, nums2, n)
}