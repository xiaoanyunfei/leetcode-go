package main

import "fmt"

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	length := len(nums)
	i := length - 1
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			j := length - 1
			for nums[j] <= nums[i-1] {
				j--
			}
			nums[i-1], nums[j] = nums[j], nums[i-1]
			reverse(nums, i)
			break
		}
	}
	if i == 0 {
		reverse(nums, 0)
	}
}

func reverse(nums []int, l int) {
	r := len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		r--
		l++
	}
}
func main() {
	// nums := []int{1, 4, 3, 8, 7, 2, 1}
	nums := []int{1, 5, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
