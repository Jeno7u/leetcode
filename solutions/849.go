package main

import "fmt"


func maxDistToClosest(seats []int) int {
	result, last := 0, -1
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if last != -1 {
				result = max(result, (i - last) / 2)
			} else {
				result = i
			}
			last = i
		}
	}
	return max(result, len(seats) - last - 1)
}

func main() {
	seats := []int{1, 0, 0, 0}
	fmt.Println(maxDistToClosest(seats))
}