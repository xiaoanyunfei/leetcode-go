package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 2))
}
