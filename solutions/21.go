package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// создаем dummy и туда добавляем меньший элемент из двух списков и сдвигаемся к след элементу
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
    head := dummy

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				dummy.Next = list1
                list1 = list1.Next
			} else {
				dummy.Next = list2
                list2 = list2.Next
			}
		} else if list1 != nil && list2 == nil {
            dummy.Next = list1
            list1 = list1.Next
        } else {
            dummy.Next = list2
            list2 = list2.Next
        }
        dummy = dummy.Next
	}
    return head.Next
}
