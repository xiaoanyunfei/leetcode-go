package main

import "fmt"

func maxArea(height []int) int {
	var maxWater, tmp int
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			tmp = height[left] * (right - left)
			left++
		} else {
			tmp = height[right] * (right - left)
			right--
		}
		if tmp > maxWater {
			maxWater = tmp
		}
	}
	return maxWater
}

func main() {
	height := []int{6, 5, 7, 3, 4}
	fmt.Println(maxArea(height))
	height = []int{1, 1}
	fmt.Println(maxArea(height))
}
