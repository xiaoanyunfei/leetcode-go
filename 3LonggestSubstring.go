package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	indexes := make(map[int32]int)
	maxLen := 0
	i := 0
	for j, v := range s {
		if index, ok := indexes[v]; ok {
			if index > i {
				i = index
			}

		}
		if j-i+1 > maxLen {
			maxLen = j - i + 1
		}
		indexes[v] = j + 1
	}
	fmt.Println(indexes)
	return maxLen
}

func main() {
	s := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring(s))
}
