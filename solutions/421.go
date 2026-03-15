package main

type TrieChildren struct {
	children [2]*TrieChildren
}

func (t *TrieChildren) InsertNum(num int) {
	binary := [32]int{}
	for i := 0; i < 32; i++ {
		if num%2 == 1 {
			binary[i] = 1
		}
		num /= 2
	}

	root := t
	for i := 31; i >= 0; i-- {
		idx := binary[i]
		if root.children[idx] == nil {
			root.children[idx] = &TrieChildren{}
		}
		root = root.children[idx]
	}
}

func (t *TrieChildren) FindXORMax(num int) int {
	binary := [32]int{}
	for i := 0; i < 32; i++ {
		if num%2 == 1 {
			binary[i] = 1
		}
		num /= 2
	}

	root := t
	result := 0
	for i := 31; i >= 0; i-- {
		var idx int
		if binary[i] == 0 {
			idx = 1
		}

		if root.children[idx] != nil {
			result = result*2 + idx
			root = root.children[idx]
		} else {
			result = result*2 + binary[i]
			root = root.children[binary[i]]
		}
	}

	return result
}

func findMaximumXOR(nums []int) int {
	TrieChildren := TrieChildren{}

	for i := range nums {
		TrieChildren.InsertNum(nums[i])
	}

	result := -1
	for i := range nums {
		result = max(result, nums[i]^TrieChildren.FindXORMax(nums[i]))
	}
	return result
}
