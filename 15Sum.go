package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i, v := range nums {
		if i > len(nums)-3 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := -v
			if nums[l]+nums[r] == sum {
				result = append(result, []int{v, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else if nums[l]+nums[r] > sum {
				r--
			} else {
				l++
			}
		}
	}
	return result
}

func main() {
	S := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(S))
	S = []int{0, 0, 0}
	fmt.Println(threeSum(S))
	S = []int{-2, 0, 1, 1, 2}
	fmt.Println(threeSum(S))
}
