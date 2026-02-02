package main

import "slices"


func diagonalSort(mat [][]int) [][]int {
    mapElements := map[int][]int{}

    for j := range mat {
        for i := range mat[0] {
            diagonalIdx := i - j
            mapElements[diagonalIdx] = append(mapElements[diagonalIdx], mat[j][i])
        }
    }

    for key, _ := range mapElements {
        slices.Sort(mapElements[key])
    }

    for j := range mat {
        for i := range mat[0] {
            diagonalIdx := i - j
            mat[j][i] = mapElements[diagonalIdx][0]
            mapElements[diagonalIdx] = mapElements[diagonalIdx][1:]
        }
    }
    return mat
}