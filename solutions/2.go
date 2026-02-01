package main


type ListNode struct {
	Val int
	Next *ListNode
}

// можно сократить решение вдвое если создать новый linked list и в него сохранять результаты, но мое решение по памяти
// O(1) за счет сохранения результатов внутри l1. В любом случае мы берем значения из l1 и l2 + extra (от предыдущих вычеслений) и
// сохраняем результат + записываем новое значение extra
func getNewVal(val int) (int, int) {
    extra := 0
    if val > 9 {
        val -= 10
        extra = 1
    }
    return val, extra
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    root := l1
    extra := 0
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			l1.Val, extra = getNewVal(l1.Val + l2.Val + extra)
            if l1.Next == nil && l2.Next != nil {
                l1.Next = l2.Next
                l2.Next = nil
            } else if l1.Next == nil && l2.Next == nil && extra == 1 {
                l1.Next = &ListNode{Val: extra}
                break
            }
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			l1.Val, extra = getNewVal(l1.Val + extra)
            if l1.Next == nil && extra == 1 {
                l1.Next = &ListNode{Val: extra}
                break
            }
            l1 = l1.Next
        }
	}

    return root
}