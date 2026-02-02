package main


func canPlaceFlowers(flowerbed []int, n int) bool {
    for i := range flowerbed {
        left := i == 0 || flowerbed[i - 1] == 0
        right := i == len(flowerbed) - 1 || flowerbed[i + 1] == 0
        if left && right && flowerbed[i] == 0 {
            flowerbed[i] = 1
            n--
        }
    }
    return n <= 0
}