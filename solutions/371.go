package main


func getSum(a int, b int) int {
    for b != 0 {
        carry := (a & b) << 1 // вычесляем биты где у нас carry и сдвигаем их влево
        a = a ^ b // объединяем либо изначально, либо с carry
        b = carry 
    }
    return a
}
