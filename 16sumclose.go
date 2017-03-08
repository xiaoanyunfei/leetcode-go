package main

import (
	"fmt"
	"sort"
)

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	for i, v := range nums {
		if i > len(nums)-3 {
			break
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := v + nums[l] + nums[r]
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
			if sum == target {
				return sum
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return result
}

func main() {
	S := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(S, 1))
}
