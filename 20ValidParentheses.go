package main

import "fmt"

type Stack struct {
	nodes []rune
	count int
}

func (s *Stack) Push(n rune) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() rune {
	var value rune
	if s.count != 0 {
		s.count--
		value = s.nodes[s.count]
	}
	return value
}

func (s *Stack) IsEmpty() bool {
	return s.count == 0
}

func isValid(s string) bool {
	store := new(Stack)
	for _, v := range s {
		if v == '[' {
			store.Push(']')
		} else if v == '{' {
			store.Push('}')
		} else if v == '(' {
			store.Push(')')
		} else if store.IsEmpty() || v != store.Pop() {
			return false
		}
	}
	return store.IsEmpty()
}

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
