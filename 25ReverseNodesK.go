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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil {
		return head
	}
	last := head
	for i := 0; i < k-1; i++ {
		last = last.Next
		if last == nil {
			return head
		}
	}
	nextHead := last.Next
	last.Next = nil
	reverse(head)
	head.Next = reverseKGroup(nextHead, k)
	return last
}

func reverse(head *ListNode) *ListNode {
	pre := new(ListNode)
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}

func main() {
	list1 := factory([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	result := reverseKGroup(list1, 3)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
