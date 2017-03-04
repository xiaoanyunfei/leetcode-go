package main

import "fmt"

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for k, v := range p {
		if v == '*' && dp[0][k-1] {
			dp[0][k+1] = true
		}
	}
	for i, a := range s {
		for j, b := range p {
			if b == '.' || b == a {
				dp[i+1][j+1] = dp[i][j]
			} else if b == '*' {
				if p[j-1] != byte(a) && p[j-1] != '.' {
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = (dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1])
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

func main() {
	fmt.Println(isMatch("aab", "c*a*b"))
}
