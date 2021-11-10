package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func List2Val(head *ListNode) []int {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}

func Val2List(values []int) *ListNode {
	n := len(values)
	if n == 0 {
		return nil
	}
	dummy := &ListNode{}
	curr := dummy
	for _, v := range values {
		node := &ListNode{Val: v}
		curr.Next = node
		curr = node
	}
	return dummy.Next
}

func Val2ListWithCycle(values []int, pos int) *ListNode {
	head := Val2List(values)
	if pos == -1 {
		return head
	}
	curr := head
	for pos > 0 {
		curr = curr.Next
		pos--
	}
	tail := curr
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = curr
	return head
}
