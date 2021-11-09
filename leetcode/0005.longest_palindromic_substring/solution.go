// 5. 最长回文子串(https://leetcode-cn.com/problems/longest-palindromic-substring/)

package leetcode

// leetcode submit region begin(Prohibit modification and deletion)

type Answer struct {
	Pos int
	Len int
}

func longestPalindrome(s string) string {
	answer := &Answer{Pos: 0, Len: 1}
	n := len(s)
	for i := 0; i < n; i++ {
		expandAroundCenter(s, i, i, answer)
		expandAroundCenter(s, i, i+1, answer)
	}
	return s[answer.Pos : answer.Pos+answer.Len]
}

func expandAroundCenter(s string, left, right int, answer *Answer) {
	n := len(s)
	for left >= 0 && right < n && s[left] == s[right] {
		left -= 1
		right += 1
	}
	length := (right - 1) - (left + 1) + 1
	if length > answer.Len {
		answer.Pos = left + 1
		answer.Len = length
	}
}

// leetcode submit region end(Prohibit modification and deletion)
