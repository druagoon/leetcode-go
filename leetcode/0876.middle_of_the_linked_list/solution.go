// 876. 链表的中间结点(https://leetcode-cn.com/problems/middle-of-the-linked-list/)

package leetcode

import (
	. "github.com/druagoon/leetcode-go/structures"
)

// leetcode submit region begin(Prohibit modification and deletion)
// Input: [1, 2, 3, 4]
// Output: [3, 4]
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// leetcode submit region end(Prohibit modification and deletion)

// Input: [1, 2, 3, 4]
// Output: [2, 3, 4]
func middleNodeLeft(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
