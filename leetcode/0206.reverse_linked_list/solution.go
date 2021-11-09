// 206. 反转链表(https://leetcode-cn.com/problems/reverse-linked-list/)

package leetcode

import (
	. "github.com/druagoon/leetcode-go/structures"
)

// leetcode submit region begin(Prohibit modification and deletion)
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

// leetcode submit region end(Prohibit modification and deletion)
