package main

import "fmt"

func generateParenthesis(n int) []string {
	var res []string
	dfs(&res, n, n, "")
	return res
}

func dfs(res *[]string, l int, r int, tmp string) {
	if l == 0 && r == 0 {
		*res = append(*res, tmp)
	}
	if l > 0 {
		dfs(res, l-1, r, tmp+"(")
	}
	if r > l {
		dfs(res, l, r-1, tmp+")")
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
