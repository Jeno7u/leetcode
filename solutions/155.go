package main

import "fmt"

type Node struct {
	Val int
	PrevMinimum int
	Prev *Node
	Next *Node
}
// ну тут хуйня конечно. Надо было просто использовать два списка. Один для stack и
// второй для хранения prevMinimum. Но и это решение неплохое, но больше строчек
type MinStack struct {
	end *Node
	minimum int
}


func Constructor() MinStack {
    return MinStack{end: nil, minimum: 0}
}


func (this *MinStack) Push(val int)  {
    if this.end == nil {
		this.end = &Node{Val: val, PrevMinimum: val}
		this.minimum = val
		return
	}

	this.end.Next = &Node{Val: val, PrevMinimum: this.minimum, Prev: this.end}
	this.minimum = min(this.minimum, val) // set new minimum
	this.end = this.end.Next
}


func (this *MinStack) Pop()  {
    if this.end.Prev == nil {
        this.minimum = 0
        this.end = nil
    } else {
	    this.minimum = this.end.PrevMinimum
        this.end = this.end.Prev
        this.end.Next = nil
    }
}


func (this *MinStack) Top() int {
    return this.end.Val
}


func (this *MinStack) GetMin() int {
    return this.minimum
}

func main() {
	stack := Constructor()

	stack.Push(2)
	stack.Push(0)
	stack.Push(3)
	stack.Push(0)
	stack.Pop()
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
}