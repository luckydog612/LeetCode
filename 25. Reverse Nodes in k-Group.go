package main

import (
	"fmt"
)

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil || head.Next == nil {
		return head
	}
	last := head
	for i := 0; i < k; i++ {
		if last == nil {
			return head
		}
		last = last.Next
	}
	newHead := reverseLN(head, last)
	head.Next = reverseKGroup(last, k)
	return newHead
}

func reverseLN(first, last *ListNode) *ListNode {
	pre := last
	for first != last {
		tmp := first.Next
		first.Next = pre
		pre = first
		first = tmp
	}
	return pre
}

func main() {
	l5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l5,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	head := &ListNode{
		Val:  1,
		Next: l2,
	}
	fmt.Printf("%+v\n", head)
	pre := reverseKGroup(head, 3)
	fmt.Printf("%+v\n", pre)
}
