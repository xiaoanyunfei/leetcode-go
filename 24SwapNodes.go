package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func factory(nodes []int) *ListNode {
	head := new(ListNode)
	pre := head
	for _, node := range nodes {
		now := new(ListNode)
		now.Val = node
		pre.Next = now
		pre = now
	}
	return head.Next
}

func swapPairs(head *ListNode) *ListNode {
	before := new(ListNode)
	before.Next = head
	record := before
	for head != nil && head.Next != nil {
		after := head.Next
		before.Next = after
		head.Next = after.Next
		after.Next = head
		before = head
		head = head.Next
	}
	return record.Next
}

func main() {
	list1 := factory([]int{1, 3, 4, 9})
	result := swapPairs(list1)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
