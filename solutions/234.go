package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeLinkedList(head *ListNode) bool {
    slow, fast := head, head

    // находим половину (slow pointer)
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // разворачиваем вторую половину
    var node *ListNode
    for slow != nil {
        tmp := slow.Next
        slow.Next = node
        node = slow
        slow = tmp
    }

    // теперь просто проверяем на палиндром
    // У нас обрезается так что, вторая часть
    // >= половине, но первая часть все еще ссылается
    // на  значение slow pointer (середины) и все короче круто
    first, second := head, node
    for second != nil {
        if first.Val != second.Val {
            return false
        }
        first = first.Next
        second = second.Next
    }
    return true
}