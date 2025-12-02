package main

import "fmt"

// diff array not used for query, it use for update with O(1)
func buildDifferenceArray(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	diff := make([]int, n)
	diff[0] = arr[0]
	for i := 1; i < n; i++ {
		diff[i] = arr[i] - arr[i-1]
	}
	return diff
}

// simple update
func update(diff []int, l, r, x, n int) []int {
	diff[l] += x
	if r+1 < n {
		diff[r+1] -= x
	}
	return diff
}

// rebuild after update
func rebuild(diff []int) []int {
	arr := make([]int, len(diff))
	arr[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		arr[i] = arr[i-1] + diff[i]
	}
	return arr
}

// ex1:
// Bài 1 – Easy: Update 1 lần trên đoạn [l..r]
// Yêu cầu bài toán
// Cho:
// Một mảng số nguyên a gồm n phần tử.
// Một thao tác duy nhất: cộng thêm giá trị x cho toàn bộ đoạn a[l..r] (0-based).
// Hãy in ra mảng sau khi thực hiện update.

func solve_ex1_diff() {
	var n, l, r, x int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&l, &r, &x)
	var diff = buildDifferenceArray(arr)
	diff_update := update(diff, l, r, x, n)
	var rs = rebuild(diff_update)
	for i := 0; i < len(rs); i++ {
		fmt.Println(rs[i])
	}
}

// ex2:
// Bài 2 – Medium: Nhiều lần update trên đoạn
// Giờ lên bài Medium như đã hứa nè.
// Đề bài
// Cho một mảng a gồm n phần tử, ban đầu tất cả đều bằng 0.
// Bạn có q thao tác, mỗi thao tác có dạng:
// Cộng thêm x vào toàn bộ đoạn a[l..r] (0-based, inclusive).
// Sau khi thực hiện xong tất cả thao tác, hãy in ra mảng cuối cùng.

func solve_ex2_diff() {
	var n, q, l, r, x int
	fmt.Scan(&n, &q)
	diff := make([]int, n)

	for i := 0; i < q; i++ {
		fmt.Scan(&l, &r, &x)
		update(diff, l, r, x, n)
	}
	arr := rebuild(diff)
	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}
}
