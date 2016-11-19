package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tag := 0
	result := &ListNode{0, nil}
	l3 := result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tag += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tag += l2.Val
			l2 = l2.Next
		}
		l3.Next = &ListNode{tag % 10, nil}
		l3 = l3.Next
		tag = tag / 10
	}
	if tag == 1 {
		l3.Next = &ListNode{tag, nil}
	}
	return result.Next
}

func factory(list []int) *ListNode {
	var previous, current *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		current = &ListNode{list[i], previous}
		previous = current
	}
	return current

}

func main() {
	a := []int{2, 4, 3}
	b := []int{5, 6, 4}
	l1 := factory(a)
	l2 := factory(b)
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}
