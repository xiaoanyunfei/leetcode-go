package main

import "fmt"

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if nums[l] == target {
		result[0] = l
	} else {
		return result
	}
	r = len(nums) - 1
	for l < r {
		m := (l+r)/2 + 1
		if nums[m] > target {
			r = m - 1
		} else {
			l = m
		}
	}
	result[1] = r
	return result
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}
