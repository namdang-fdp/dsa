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

// -------------- More Exercises----------------
// Cho mảng số nguyên A[] gồm N phần tử, mảng cộng dồn của A[] là mảng F[] với
// tính chất F[i] lưu tổng các phần tử từ chỉ số 0 tới chỉ số i của mảng A[]. Bạn hãy  xây dựng mảng cộng dồn F[]
func solve_b1(arr []int, n int) []int {
	prefix := make([]int, n)
	prefix[0] = arr[0]
	for i := 1; i < n; i++ {
		prefix[i] = arr[i] + prefix[i-1]
	}
	return prefix
}

// Cho mảng số nguyên A[] gồm N phần tử, có Q truy vấn, mỗi truy vấn là 2 số L, R  bạn hãy tính tổng các số từ chỉ số L tới chỉ số R của mảng.
func solve_b2() {
	var n, q int
	fmt.Scan(&n, &q)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	prefix := solve_b1(arr, n)
	for i := 0; i < q; i++ {
		var left, right int
		fmt.Scan(&left, &right)
		if left == 0 {
			fmt.Println(prefix[right])
		} else {
			fmt.Println(prefix[right] - prefix[left-1])
		}
	}
}

// Tèo hiện tại đã bỏ làm lập trình viên và trở về về quê trồng rau nuôi cá, anh ta có một mảnh vườn hình chữ nhật có kích thước NxM. Anh ta chia vườn của mình  thành NXM ô vuông và trồng vào đó một cây cà rốt, tới vụ thu hoạch có những  cây cà rốt bị chết và có những cây cà rốt có củ, anh ta muốn biết với mỗi mảnh  vườn hình chữ nhật bắt đầu từ hàng x1 tới hàng x2 và từ cột y1 tới cột y2 thì số  cà rốt thu hoạch được là bao nhiêu.
// Biết rằng mảnh vườn được mô tả bởi một ma trận nhị phân, 0 tương ứng với cây  cà rốt chất và 1 tương ứng với cây cà rốt có củ.

func build_prefix_matrix(arr [][]int, n, m int) [][]int {
	prefix_matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		prefix_matrix[i] = make([]int, m)
	}
	for i := 0; i < len(arr); i++ {
		prefix_matrix[i][0] = arr[i][0]
		for j := 1; j < len(arr[i]); j++ {
			prefix_matrix[i][j] = arr[i][j] + prefix_matrix[i][j-1]
		}
	}
	return prefix_matrix
}

// 1 2 2 2 2 3 4 5
// 1 1 1 1 2 3 4 5
// 0 1 2 3 3 3 4 4
func solve_b3() {
	var n, m int
	fmt.Scan(&n, &m)
	matrix_garden := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix_garden[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix_garden[i][j])
		}
	}
	prefix_garden := build_prefix_matrix(matrix_garden, n, m)
	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var x1, x2, y1, y2 int
		fmt.Scan(&x1, &x2, &y1, &y2)
		if y1 == 1 {
			fmt.Print(prefix_garden[x1-1][y2-1] + prefix_garden[x2-1][y2-1])
		} else if x1 == x2 {
			fmt.Println(prefix_garden[x1-1][y2-1] - prefix_garden[x1-1][y1-2])
		} else {
			fmt.Println((prefix_garden[x1-1][y2-1] - prefix_garden[x1-1][y1-2]) + (prefix_garden[x2-1][y2-1] - prefix_garden[x2-1][y1-2]))
		}
	}
}
