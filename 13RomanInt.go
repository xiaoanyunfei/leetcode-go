package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	var sum int
	if strings.Contains(s, "IV") {
		sum -= 2
	}
	if strings.Contains(s, "IX") {
		sum -= 2
	}
	if strings.Contains(s, "XL") {
		sum -= 20
	}
	if strings.Contains(s, "XC") {
		sum -= 20
	}
	if strings.Contains(s, "CD") {
		sum -= 200
	}
	if strings.Contains(s, "CM") {
		sum -= 200
	}
	for _, v := range s {
		switch v {
		case 'M':
			sum += 1000
		case 'D':
			sum += 500
		case 'C':
			sum += 100
		case 'L':
			sum += 50
		case 'X':
			sum += 10
		case 'V':
			sum += 5
		case 'I':
			sum += 1
		}
	}
	return sum
}

func main() {
	s := "MXXIV"
	fmt.Println(romanToInt(s))

}
