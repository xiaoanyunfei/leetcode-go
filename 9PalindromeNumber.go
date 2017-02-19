package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else {
		tmp := x
		y := 0
		for x > 0 {
			y = y*10 + x%10
			x = x / 10
		}
		if y == tmp {
			return true
		} else {
			return false
		}
	}
}

func main() {
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(123321))
	fmt.Println(isPalindrome(154451))
}
