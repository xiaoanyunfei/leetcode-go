package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := range haystack {
		for j := range needle {
			if needle[j] != haystack[i+j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
			if i+j == len(haystack)-1 {
				return -1
			}
		}
	}
	return -1
}

func main() {
	haystack := "abcdefcde"
	needle := "cde"
	fmt.Println(strStr(haystack, needle))
}
