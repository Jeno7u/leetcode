package main


func setZeroes(matrix [][]int)  {
    // храним параметр зануления строки и столбца в первой
    // строке и столбце. Но чтобы случайно не занулить лишнее
    // (а такое будет если у нас есть 0 в первой строке или стобце)
    // мы отдельно проверяем на наличие нулей в этих случаях и 
    // затем заполняем нулями на основе наших параметров. И потом
    // только делаем заполнение наших первых столбцов и стсрочек (если надо)
    rowZero := false
    columnZero := false
    for i := range matrix[0] {
        if matrix[0][i] == 0 {
            rowZero = true
            break
        }
    }

    for j := range matrix {
        if matrix[j][0] == 0 {
            columnZero = true
            break
        }
    }

    for j := 1; j < len(matrix); j++ {
        for i := 1; i < len(matrix[0]); i++ {
            if matrix[j][i] == 0 { 
                matrix[0][i] = 0
                matrix[j][0] = 0
            }
        }
    }

    for i := 1; i < len(matrix[0]); i++ {
        if matrix[0][i] == 0 {
            for j := 1; j < len(matrix); j++ {
                matrix[j][i] = 0
            }
        }
    }

    for j := 1; j < len(matrix); j++ {
        if matrix[j][0] == 0 {
            for i := 1; i < len(matrix[0]); i++ {
                matrix[j][i] = 0
            }
        }
    }

    if rowZero {
        for i := range matrix[0] {
            matrix[0][i] = 0
        }
    }

    if columnZero {
        for j := range matrix {
            matrix[j][0] = 0
        }
    }
}