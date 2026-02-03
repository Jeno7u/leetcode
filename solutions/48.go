package main


func rotate(matrix [][]int)  {
    l, r := 0, len(matrix) - 1
    for l < r {
        for i := 0; i < (r - l); i++ {
            top, bottom := l, r

            // сохраняем значение
            topLeft := matrix[top][l + i]

            // снизу слева -> верхне левое
            matrix[top][l + i] = matrix[bottom - i][l]

            // нижне правое -> нижне левое
            matrix[bottom - i][l] = matrix[bottom][r - i]

            // верхне правое -> нижне правое
            matrix[bottom][r - i] = matrix[top + i][r]
            
            // верхне левое -> верхне правое
            matrix[top + i][r] = topLeft
        }
        l++
        r--
    }
}