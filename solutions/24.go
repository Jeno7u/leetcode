package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// создаем два поинтера (left и right). Меняем их местами и соединяем с предыдущим node чтобы не было циклов
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummy := head.Next
    var prev *ListNode = nil
    left, right := head, head.Next
    for left != nil && right != nil {
        // swap nodes
        tmp := right.Next
        right.Next = left
        left.Next = tmp
        // если prev существует, мы соединяем prev и текущий left
        if prev != nil {
            prev.Next = right
        } 
        prev = left

        if left.Next == nil || left.Next.Next == nil {
            break
        }
        // идем на следующую итерацию
        left = left.Next
        right = left.Next
    }
    return dummy
}

