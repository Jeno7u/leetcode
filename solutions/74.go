package main

func searchMatrix(matrix [][]int, target int) bool {
    // find row
    l, r := 0, len(matrix)
    for l < r {
        mid := (l + r) / 2
        if matrix[mid][0] >= target {
            r = mid
        } else {
            l = mid + 1
        }
    }
    rowIdx := l 
    if rowIdx == 0 && matrix[0][0] > target {
        return false
    }
    if rowIdx != len(matrix) && matrix[rowIdx][0] == target {
        return true
    }
    rowIdx-- // l это индекс куда должен быть вставлен target, так что - 1

    l, r = 0, len(matrix[0])
    for l < r {
        mid := (l + r) / 2
        if matrix[rowIdx][mid] >= target {
            r = mid
        } else {
            l = mid + 1
        }
    }

    columnIdx := l
    if columnIdx == len(matrix[0]) {
        return false
    }
    if matrix[rowIdx][columnIdx] == target {
        return true
    }
    return false    
}