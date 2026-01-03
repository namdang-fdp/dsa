package main

import (
	"fmt"
	"sort"
)

// Đề:
// Cho mảng đã sort tăng dần nums và target.
// Hãy trả lời xem có tồn tại cặp (i, j) sao cho nums[i] + nums[j] = target hay không.
// Trả về true/false là đủ.

func twoSumExistsSorted(nums []int, target int) bool {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}

// Đề:
// Cho mảng chưa sort nums và target.
// Hãy trả về tất cả cặp (a, b) sao cho:
// a + b = target
// Không có cặp trùng (ví dụ [1,2] chỉ xuất hiện một lần).
func twoSumAllUnique(nums []int, target int) [][]int {
	// -1 -1 0 1 1 2 2 3 4 4
	sort.Ints(nums)
	freq := make(map[int]int)
	res := [][]int{}
	left := 0
	right := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		sum := nums[left] + nums[right]
		if sum == target && freq[nums[left]] == 0 && freq[nums[right]] == 0 {
			res = append(res, []int{nums[left], nums[right]})
			freq[nums[left]]++
			freq[nums[right]]++
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			right--
			left++
		}
	}
	return res
}

func main() {
	var n, target int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&target)
	fmt.Print(twoSumAllUnique(arr, target))
}
