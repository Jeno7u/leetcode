package main

import (
	"math"
)

// это пиздец. идея в том чтобы найти два window в первом и втором списке, то при их объединении получится корректная половина объединенных списков.
// в таком случае мы сможем просто определеить значения в середине, как последнее в нашем window и последнее + 1 (ну там на самом деле max() и min() еще)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    A, B := nums1, nums2
	total := len(A) + len(B)
	half := total / 2

	// so A will be smaller than B
	if len(A) > len(B) {
		A, B = B, A
	}

	l, r := -1, len(A) - 1
	for l <= r {
		i := (l + r) / 2 // len of window in A
		j := half - i - 2// len of window in B

		ALeft := math.Inf(-1)
		if i >= 0  {
			ALeft = float64(A[i])
		}
		ARight := math.Inf(1)
		if i + 1 < len(A) {
			ARight = float64(A[i + 1])
		}
		BLeft := math.Inf(-1)
		if j >= 0 {
			BLeft = float64(B[j])
		}
		BRight := math.Inf(1)
		if j + 1 < len(B) {
			BRight = float64(B[j + 1])
		}

		if ALeft <= BRight && BLeft <= ARight {
			if total % 2 == 0 {
				return (max(ALeft, BLeft) + min(ARight, BRight)) / 2
			}
			return min(ARight, BRight)
		} else if ALeft > BRight{
			r = i - 1
		} else {
			l = i + 1
		}
	}
	return 0
}
