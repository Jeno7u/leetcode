package main


func judgeCircle(moves string) bool {
    dx, dy := 0, 0
    for _, move := range moves {
        switch move {
            case 'R':
                dx++
            case 'L':
                dx--
            case 'U':
                dy++
            case 'D':
                dy--
        }
    }
    
    return dx == 0 && dy == 0
}