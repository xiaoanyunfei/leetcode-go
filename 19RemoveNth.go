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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	start := new(ListNode)
	start.Next = head
	fast := start
	slow := start
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return start.Next
}

func main() {
	nodes := []int{1, 2, 3, 4, 5}
	head := factory(nodes)
	result := removeNthFromEnd(head, 2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
