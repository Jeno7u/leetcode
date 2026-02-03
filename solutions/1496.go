package main


func isPathCrossing(path string) bool {
    visited := map[[2]int]struct{}{}

    i, j := 0, 0
    visited[[2]int{i, j}] = struct{}{}
    for idx := range path {
        switch path[idx] {
            case 'N':
                j++
            case 'S':
                j--
            case 'E':
                i++
            case 'W':
                i--
        }
        if _, ok := visited[[2]int{i, j}]; ok {
            return true
        }
        visited[[2]int{i, j}] = struct{}{}
    }
    return false
}