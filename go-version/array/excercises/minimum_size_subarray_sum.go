// this is leetcode 209
package main

func minSubArrayLen(target int, nums []int) int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]
	for i := 0; i < len(nums); i++ {
		prefix[i] = nums[i] + prefix[i-1]
	}
	if prefix[len(nums)-1] < target {
		return 0
	}
	return 0
}
