package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	currentList := list1

	if currentList.Next != nil {
		fmt.Println(currentList.Val)
	}
	return list1
}
