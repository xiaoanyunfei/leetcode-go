package main

import "fmt"

func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result > 0x7fffffff || result < -0x80000000 {
		result = 0
	}
	return result
}

func main() {
	x := 1534236469
	fmt.Println(reverse(x))
	x = -123
	fmt.Println(reverse(x))
}
