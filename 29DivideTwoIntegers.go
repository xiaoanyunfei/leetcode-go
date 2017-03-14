package main

import (
	"fmt"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func divide(dividend int, divisor int) int {
	res := 0
	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}
	dividend = abs(dividend)
	divisor = abs(divisor)
	for dividend >= divisor {
		tmp := divisor
		i := 1
		for dividend >= tmp<<1 {
			tmp <<= 1
			i <<= 1
		}
		res += i
		dividend -= tmp
	}
	res = sign * res
	return min(max(-2147483648, res), 2147483647)
}

func main() {
	fmt.Println(divide(1000, 11))
}
