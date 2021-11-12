// 234. 回文链表(https://leetcode-cn.com/problems/palindrome-linked-list/)

package leetcode

import (
	. "github.com/druagoon/leetcode-go/structures"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var curr, prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		curr = slow
		slow = slow.Next
		fast = fast.Next.Next
		curr.Next = prev
		prev = curr
	}
	if fast != nil {
		slow = slow.Next
	}
	for curr != nil && slow != nil {
		if curr.Val != slow.Val {
			return false
		}
		curr = curr.Next
		slow = slow.Next
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)
