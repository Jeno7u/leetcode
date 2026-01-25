package main


type ListNode struct {
	Val int
	Next *ListNode
}

// можно сократить решение вдвое если создать новый linked list и в него сохранять результаты, но мое решение по памяти
// O(1) за счет сохранения результатов внутри l1. В любом случае мы берем значения из l1 и l2 + extra (от предыдущих вычеслений) и
// сохраняем результат + записываем новое значение extra
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    root := l1
    extra := 0
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			newVal := l1.Val + l2.Val + extra
			if newVal > 9 {
				newVal -= 10
				extra = 1
			} else {
				extra = 0
			}
			l1.Val = newVal
			
			if l1.Next == nil && l2.Next != nil {
				l1.Next = &ListNode{Val: 0, Next: nil}
			}
            if l1.Next == nil && l2.Next == nil && extra == 1 {
                l1.Next = &ListNode{Val: 1, Next: nil}
                extra = 0
            }

			l1 = l1.Next
			l2 = l2.Next
		}
		if l1 != nil && l2 == nil {
			newVal := l1.Val + extra
			if newVal == 10 {
				newVal = 0
				extra = 1
			} else {
				extra = 0
			}
			l1.Val = newVal

            if l1.Next == nil && extra == 1 {
                l1.Next = &ListNode{Val: 1, Next: nil}
                extra = 0
            }
			l1 = l1.Next
		}
	}

    return root
}
