package main

import "container/heap"

// если создать minHeap с размером k и пройтись по nums убирая значения минимальное, если нашли
// значение которое больше.
func findKthLargest(nums []int, k int) int {
    h := IntHeap(nums[:k])
    heap.Init(&h)
    for _, v := range nums[k:] {
        if v > h[0] {
            heap.Pop(&h)
            heap.Push(&h, v)
        }
    }
    return h[0]
}

// int min heap via https://pkg.go.dev/container/heap#example-package-IntHeap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *IntHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}