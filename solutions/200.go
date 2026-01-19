package main

import "fmt"

func getNeighbours(grid [][]byte, i, j int) [][]int{
	result := [][]int{}
	di := []int{-1, 0, 0, 1}
	dj := []int{0, -1, 1, 0}
	for idx := range di {
		new_i := i + di[idx]
		new_j := j + dj[idx]
		if 0 <= new_i && new_i < len(grid[0]) && 0 <= new_j && new_j < len(grid) {
			if grid[new_j][new_i] == '1' {
				result = append(result, []int{new_i, new_j})
			}
		}
	}
	return result
}

func clearIsland(grid [][]byte, i, j int) {
	if grid[j][i] == '0' {
		return
	}
	grid[j][i] = '0' // чтобы не посещать те же клетки два раза
	neighbours := getNeighbours(grid, i, j)
	for _, neighbour := range neighbours {
		neighbour_i := neighbour[0]
		neighbour_j := neighbour[1]
		if grid[neighbour_j][neighbour_i] == '0' {
			continue
		}
		clearIsland(grid, neighbour_i, neighbour_j)
	}
}
// когда находим сушу, то мы исследуем весь остров и его заменяем на '0'.
func numIslands(grid [][]byte) int {
	result := 0
    for j := range grid {
		for i := range grid[0] {
			if grid[j][i] == '1' {
				clearIsland(grid, i, j)
				result++
			}
		}
	}
	return result
}


func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}