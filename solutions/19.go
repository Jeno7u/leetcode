package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// вместо того чтобы находить последний элемент, потом предпоследний и так далее,
// можно один поинтер сместить на n и создать второй поинтер равный head. Тогда
// расстояние между ними будет n и если сдвигать два поинтера пока правый не дойдет до конца,
// то левый поинтер окажется как раз на n нодов от конца
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    right := head
    for i := 0; i < n; i++ {
        right = right.Next
    }
    left := head

    if right == nil {
        return head.Next
    }

    for right.Next != nil {
        left = left.Next
        right = right.Next
    }
    left.Next = left.Next.Next

    return head
}


