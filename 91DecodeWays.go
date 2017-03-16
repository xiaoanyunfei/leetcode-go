package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	return dfs(s, 0)
}

func dfs(s string, pos int) int {
	if pos > len(s)-1 {
		return 1
	}
	if s[pos] == '0' {
		return 0
	}
	if pos == len(s)-1 {
		return 1
	}
	if s[pos] == '1' {
		return dfs(s, pos+1) + dfs(s, pos+2)
	} else if s[pos] == '2' && s[pos+1] <= '6' {
		return dfs(s, pos+1) + dfs(s, pos+2)
	} else {
		return dfs(s, pos+1)
	}
}

func main() {
	s := "26"
	fmt.Println(numDecodings(s))
}
