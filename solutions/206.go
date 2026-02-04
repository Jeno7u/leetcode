package main

// вот самый короткий вариант
func reverseList(head *ListNode) *ListNode {
    var node *ListNode
    for head != nil {
        tmp := head.Next
        head.Next = node
        node = head
        head = tmp
    }
    return node
}