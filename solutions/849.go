package main


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
