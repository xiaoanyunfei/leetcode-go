package main

import "fmt"

func findSubstring(s string, words []string) []int {
	counts := make(map[string]int)
	var result []int
	if len(words) == 0 {
		return result
	}
	for _, word := range words {
		if value, ok := counts[word]; ok {
			counts[word] = value + 1
		} else {
			counts[word] = 1
		}
	}
	m := len(words)
	n := len(words[0])
	for i := 0; i <= len(s)-m*n; i++ {
		tmp := make(map[string]int)
		for k, v := range counts {
			tmp[k] = v
		}
		for j := 0; j < m; j++ {
			word := s[i+j*n : i+j*n+n]
			if value, ok := tmp[word]; ok {
				if value == 1 {
					delete(tmp, word)
				} else {
					tmp[word] = value - 1
				}
			} else {
				break
			}
		}
		if len(tmp) == 0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	// s := "barfoothefoobarman"
	// words := []string{"foo", "bar"}
	// fmt.Println(findSubstring(s, words))
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring(s, words))
}
