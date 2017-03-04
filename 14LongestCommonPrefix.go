package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	for _, v := range strs {
		for !strings.HasPrefix(v, pre) {
			pre = pre[:len(pre)-1]
		}

	}
	return pre
}

func main() {
	strs := []string{"abcde", "abc", "ab"}
	fmt.Println(longestCommonPrefix(strs))
}
