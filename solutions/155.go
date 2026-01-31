package main

// type Node struct {
// 	Val int
// 	PrevMinimum int
// 	Prev *Node
// 	Next *Node
// }
// ну тут хуйня конечно. Надо было просто использовать два списка. Один для stack и
// второй для хранения prevMinimum. Но и это решение неплохое, но больше строчек
// type MinStack struct {
// 	end *Node
// 	minimum int
// }

// func AnotherConstructor() MinStack {
//     return MinStack{end: nil, minimum: 0}
// }

// func (this *MinStack) Push(val int)  {
//     if this.end == nil {
// 		this.end = &Node{Val: val, PrevMinimum: val}
// 		this.minimum = val
// 		return
// 	}

// 	this.end.Next = &Node{Val: val, PrevMinimum: this.minimum, Prev: this.end}
// 	this.minimum = min(this.minimum, val) // set new minimum
// 	this.end = this.end.Next
// }

// func (this *MinStack) Pop()  {
//     if this.end.Prev == nil {
//         this.minimum = 0
//         this.end = nil
//     } else {
// 	    this.minimum = this.end.PrevMinimum
//         this.end = this.end.Prev
//         this.end.Next = nil
//     }
// }

// func (this *MinStack) Top() int {
//     return this.end.Val
// }

// func (this *MinStack) GetMin() int {
//     return this.minimum
// }

// читать выше
type MinStack struct {
    minimum int
    keys []int
    prevMinimum []int
}


func AnotherConstructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    this.keys = append(this.keys, val)

    if len(this.keys) == 1 {
        this.minimum = val
        this.prevMinimum = append(this.prevMinimum, 0)
    } else {
        this.prevMinimum = append(this.prevMinimum, this.minimum)
        this.minimum = min(this.minimum, val)
    }
}


func (this *MinStack) Pop()  {
    this.minimum = this.prevMinimum[len(this.prevMinimum) - 1]
    this.keys = this.keys[:len(this.keys) - 1]
    this.prevMinimum = this.prevMinimum[:len(this.prevMinimum) - 1]
}


func (this *MinStack) Top() int {
    return this.keys[len(this.keys) - 1]
}


func (this *MinStack) GetMin() int {
    return this.minimum
}
