package main

// every element go in/out window 1 time
// fixed-size window: increase by update, not re-calculate

// MaxSumFixedWindow minh họa:
// - Mỗi phần tử vào/ra cửa sổ tối đa 1 lần
// - Cửa sổ cố định độ dài k, update sum bằng cách cộng/trừ gia tăng.
//
// nums: mảng int
// k: độ dài cửa sổ

func MaxSumFixedWindow(nums []int, k int) int {
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	maxSum := windowSum
	for right := k; right < len(nums); right++ {
		left := right - k
		windowSum += nums[right]
		windowSum -= nums[left]

		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

// LongestSubarraySumLE minh họa variable window với mảng không âm.
// Ta luôn giữ sum của cửa sổ [left..right] <= target bằng cách co left.
// Longest subarray có tổng ≤ S
func LongestSubarraySumLE(nums []int, target int) int {
	left := 0
	sum := 0
	ans := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum > target && left < right {
			sum -= nums[left]
			left++
		}
		windowLen := right - left + 1
		if windowLen > ans {
			ans = windowLen
		}
	}
	return ans
}

// LongestSubstringAtMost2 minh họa:
// - Sliding window trên string
// - Sử dụng map[byte]int để quản lý tần suất ký tự
// - Condition: mỗi ký tự xuất hiện duy nhất 1 lần.
func LongestSubstringNoRepeat(s string) (string, int) {
	freq := make(map[byte]int)
	left := 0
	maxLen := 0
	start := 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		freq[c]++

		for freq[c] > 1 {
			leftChar := s[left]
			freq[leftChar]--
			left++
		}

		windowLen := right - left + 1
		if windowLen > maxLen {
			maxLen = windowLen
			start = left
		}
	}
	substring := s[start : start+maxLen]
	return substring, maxLen
}
