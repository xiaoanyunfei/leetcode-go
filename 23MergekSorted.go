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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeTwoLists(result, lists[i])
	}
	return result
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	start := new(ListNode)
	result := start
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			start.Next = l1
			start = start.Next
			l1 = l1.Next
		} else {
			start.Next = l2
			start = start.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		start.Next = l1
	}
	if l2 != nil {
		start.Next = l2
	}
	return result.Next
}

func main() {
	list1 := []int{1, 3, 4, 9}
	list2 := []int{2, 5, 10}
	list3 := []int{4, 6, 8}
	var lists []*ListNode
	lists = append(lists, factory(list1))
	lists = append(lists, factory(list2))
	lists = append(lists, factory(list3))
	result := mergeKLists(lists)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}
