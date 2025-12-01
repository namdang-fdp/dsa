package main

import "fmt"

// prefix[i] = a[0] + a[1] +....+a[i]
// suitable for sum in the subarray
// 1. Build the prefix sum data structure
func buildPrefixSum(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	prefix := make([]int, len(arr))
	prefix[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefix[i] = arr[i] + prefix[i-1]
	}
	return prefix
}

// // Ex1: Range sum query
// Cho một mảng số nguyên arr gồm N phần tử.
// Có Q truy vấn, mỗi truy vấn cho hai chỉ số L và R (theo index 0-based) và yêu cầu:
// Tính tổng các phần tử từ arr[L] đến arr[R] (bao gồm cả L và R).
func rangeSum(prefix []int, l, r int) int {
	if l == 0 {
		return prefix[r]
	}
	return prefix[r] - prefix[l-1]
}

func solve_ex1() {
	var n, q int
	fmt.Scan(&n, &q)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	prefix := buildPrefixSum(arr)
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		ans := rangeSum(prefix, l, r)
		fmt.Println(ans)
	}
}

// ex2:
// Đếm số đoạn con có tổng bằng 0
// Mô tả

// Cho một mảng số nguyên arr gồm N phần tử.
// Hãy đếm số lượng đoạn con liên tiếp (subarray) có tổng bằng 0.

// Đoạn con là dãy các phần tử liên tiếp: arr[l], arr[l+1], ..., arr[r].
// Important: if subarray[l,r] have sum = 0 --> prefix[r] = prefix[l - 1]
func solve_ex2() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var k int = 0
	if n == 1 && arr[0] == 0 {
		k++
		fmt.Println(k)
		return
	}
	prefix := buildPrefixSum(arr)
	for l := 0; l < len(prefix); l++ {
		for r := l + 1; r < len(prefix); r++ {
			if l == 0 && prefix[r] == 0 {
				k++
			}
			if l != 0 && prefix[l-1] == prefix[r] {
				k++
			}
		}
	}
	fmt.Println(k)
}

func solve_ex2_optimize() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	count := 0
	prefix := 0
	freq := make(map[int]int)
	freq[0] = 1
	for i := 0; i < n; i++ {
		prefix += arr[i]
		if c, exist := freq[prefix]; exist {
			count += c
		}
		freq[prefix]++
	}
	fmt.Println(count)
}
