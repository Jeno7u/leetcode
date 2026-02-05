package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }

    // найти центр (slow pointer)
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    second := slow.Next
    slow.Next = nil

    // развернуть linked list с центра до конца
    var node *ListNode
    for second != nil {
        tmp := second.Next
        second.Next = node
        node = second
        second = tmp
    }

    // чередуя вставляем ноды двух списков
    head1, head2 := head, node
    for head2 != nil {
        if head1 == head2 {
            return
        }
        tmp1 := head1.Next
        tmp2 := head2.Next
        head1.Next = head2
        head2.Next=tmp1

        head1, head2 = tmp1, tmp2
    }
}