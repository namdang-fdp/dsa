// this is leetcode 15
package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	left := 0
	right := len(nums) - 1
	target := left + 1
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		sum := nums[left] + nums[right] + nums[target]
		if sum == 0 && left < right && freq[nums[left]] == 0 && freq[nums[right]] == 0 {
			res = append(res, []int{nums[left], nums[right], nums[target]})
			freq[nums[left]]++
			freq[nums[right]]++
			left++
			right--
		} else if sum > 0 {
			right--
		} else {
			left++
		}
	}
	return res
}
