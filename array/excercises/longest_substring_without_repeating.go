// this is leetcode 003 problem
package main

// abcabcbb
// pwwkew
func lengthOfLongestSubstring(s string) int {
	freq := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		freq[c]++

		for freq[c] > 1 {
			leftChar := s[left]
			freq[leftChar]--
			left++
		}
		subStringLen := right - left + 1
		if subStringLen > maxLen {
			maxLen = subStringLen
		}
	}
	return maxLen
}
