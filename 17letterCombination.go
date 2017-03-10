package main

import "fmt"

func letterCombinations(digits string) []string {
	mapping := []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{""}
	for _, d := range digits {
		var tmp []string
		index := d - '0'
		for i := 0; i < len(result); i++ {
			for j := 0; j < len(mapping[index]); j++ {
				tmp = append(tmp, result[i]+string(mapping[index][j]))
			}
		}
		result = tmp
	}
	return result
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
	digits = ""
	fmt.Println(letterCombinations(digits))
}
