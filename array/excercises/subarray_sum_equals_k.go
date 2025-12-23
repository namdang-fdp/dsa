// This is leetcode problem 560
package main

import "fmt"

// Important: prefix[r] = prefix[l - 1] + k
func subarraySum(nums []int, k int) int {
	prefix := 0
	count := 0
	freq := make(map[int]int)
	freq[0] = 1
	for i := 0; i < len(nums); i++ {
		prefix += nums[i]
		if c, exist := freq[prefix-k]; exist {
			count += c
		}
		freq[prefix]++
	}
	return count
}

func solve_560() {
	var n, k int
	fmt.Scan(&n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var count int = subarraySum(arr, k)
	fmt.Println(count)
}
