// 143. 重排链表(https://leetcode-cn.com/problems/reorder-list/)

package leetcode

import (
	. "github.com/druagoon/leetcode-go/structures"
)

// leetcode submit region begin(Prohibit modification and deletion)
func middleNodeLeft(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	curr := head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func mergeList(left, right *ListNode) {
	p1, p2 := left, right
	for p1 != nil && p2 != nil {
		p1Next := p1.Next
		p2Next := p2.Next
		p1.Next = p2
		p1 = p1Next
		p2.Next = p1
		p2 = p2Next
	}
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	left := head
	mid := middleNodeLeft(left)
	right := mid.Next
	mid.Next = nil
	right = reverseList(right)
	mergeList(left, right)
}

// leetcode submit region end(Prohibit modification and deletion)
