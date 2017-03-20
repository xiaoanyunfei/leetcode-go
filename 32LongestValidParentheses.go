package main

import "fmt"

type Stack struct {
	nodes []int
	count int
}

func (s *Stack) Push(n int) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Peek() int {
	var value int
	if s.count != 0 {
		value = s.nodes[s.count-1]
	}
	return value
}

func (s *Stack) Pop() int {
	var value int
	if s.count != 0 {
		s.count--
		value = s.nodes[s.count]
	}
	return value
}

func (s *Stack) IsEmpty() bool {
	return s.count == 0
}

func longestValidParentheses(s string) int {
	longest := 0
	index := -1
	store := new(Stack)
	for i, v := range s {
		if v == '(' {
			store.Push(i)
		} else {
			if store.IsEmpty() {
				index = i
			} else {
				store.Pop()
				if store.IsEmpty() {
					tmp := i - index
					if tmp > longest {
						longest = tmp
					}
				} else {
					tmp := i - store.Peek()
					if tmp > longest {
						longest = tmp
					}

				}
			}
		}
	}
	return longest
}

func main() {
	s := "()"
	fmt.Println(longestValidParentheses(s))
}
