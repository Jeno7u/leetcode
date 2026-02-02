package main


func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    var prev *ListNode
    next := head.Next
    for head != nil {
        head.Next = prev
        if next == nil {
            break
        }
        prev, head, next = head, next, next.Next
    }
    return head
}