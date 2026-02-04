type MyQueue struct {
    stack1 []int
    stack2 []int
}


func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    this.stack1 = append(this.stack1, x)
}


func (this *MyQueue) Pop() int {
    if len(this.stack2) == 0 {
        for len(this.stack1) != 0 {
            this.stack2 = append(this.stack2, this.stack1[len(this.stack1) - 1])
            this.stack1 = this.stack1[:len(this.stack1) - 1]
        }
    }

    value := this.stack2[len(this.stack2) - 1]
    this.stack2 = this.stack2[:len(this.stack2) - 1]
    return value
}


func (this *MyQueue) Peek() int {
    if len(this.stack2) == 0 {
        for len(this.stack1) != 0 {
            this.stack2 = append(this.stack2, this.stack1[len(this.stack1) - 1])
            this.stack1 = this.stack1[:len(this.stack1) - 1]
        }
    }

    return this.stack2[len(this.stack2) - 1]
}


func (this *MyQueue) Empty() bool {
    if len(this.stack1) == 0 && len(this.stack2) == 0 {
        return true
    }
    return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */