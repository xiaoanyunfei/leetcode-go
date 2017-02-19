package main

import "fmt"

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func myAtoi(str string) int {
	sign := 1
	index := 0
	var result int = 0
	for i, s := range str {
		if s == ' ' {
			index = i + 1
			continue
		} else if s == '+' {
			index = i + 1
			break
		} else if s == '-' {
			index = i + 1
			sign = -1
			break
		}
	}
	for _, s := range str[index:] {
		if s >= '0' && s <= '9' {
			result = result*10 + int(s-'0')
		} else {
			break
		}
	}
	result = result * sign
	if result > MaxInt {
		result = MaxInt
	} else if result < MinInt {
		result = MinInt
	}
	return result
}

func main() {
	str := "2147483648"
	fmt.Println(myAtoi(str))
	str = "  010"
	fmt.Println(myAtoi(str))
	str = "+-123"
	fmt.Println(myAtoi(str))
}
