package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	maxlen := 0
	result := ""
	for i := 0; i < len(s); i++ {
		a := i
		b := i + 1
		for a >= 0 && b < len(s) && s[a] == s[b] {
			a--
			b++
		}
		if b-a-1 > maxlen {
			maxlen = b - a - 1
			result = s[a+1 : b]
		}

		a = i - 1
		b = i + 1
		for a >= 0 && b < len(s) && s[a] == s[b] {
			a--
			b++
		}
		if b-a-1 > maxlen {
			maxlen = b - a - 1
			result = s[a+1 : b]
		}

	}

	return result
}

func main() {
	s := "babab"
	fmt.Println(longestPalindrome(s))
	s = "ccc"
	fmt.Println(longestPalindrome(s))

}
