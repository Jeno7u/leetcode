package main


func reverseBits(n int) int {
    reversed := 0
    bitCount := 0
    for n > 0 {
        reversed = (reversed << 1) + (n % 2)
        n /= 2
        bitCount++
    }
    return reversed << (32 - bitCount)
}