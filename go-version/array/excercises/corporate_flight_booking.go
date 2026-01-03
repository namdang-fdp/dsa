// this is the leetcode 1109 problem
package main

func solve_1109(bookings [][]int, n int) []int {
	diff := make([]int, n)
	for i := 0; i < len(bookings); i++ {
		diff[bookings[i][0]-1] += bookings[i][2]
		if bookings[i][1]-1 < n-1 {
			diff[bookings[i][1]] -= bookings[i][2]
		}
	}
	arr := make([]int, n)
	arr[0] = diff[0]
	for i := 1; i < n; i++ {
		arr[i] = diff[i] + arr[i-1]
	}
	return arr
}
