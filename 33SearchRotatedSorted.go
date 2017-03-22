package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] >= nums[l] {
			if nums[l] <= target && nums[m] > target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[r] >= target && nums[m] < target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 1, 2, 3}
	fmt.Println(search(nums, 8))
}
