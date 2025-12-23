package main

import "fmt"

func Solution(A []int) int {
	freq := make(map[int]int)
	for i := 0; i < len(A); i++ {
		freq[A[i]]++
	}
	// 1 2 ; 3 1; 6 1; 4 1; 2 1
	for i := 0; i < len(A); i++ {
		if freq[A[i]] > 0 && freq[A[i]+1] == 0 {
			return A[i] + 1
		}
	}
	return 1
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Print("here")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
	}
	fmt.Print(Solution(arr))
}
