package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur, next *ListNode

	prev = nil
	cur = head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head
	var prev, nthFromEnd *ListNode
	count := 0

	for ; curr != nil; curr = curr.Next {
		count++
		if n == count {
			nthFromEnd = head
		}
		if count > n {
			prev = nthFromEnd
			nthFromEnd = nthFromEnd.Next
		}
	}
	if count == n {
		head = head.Next
		nthFromEnd.Next = nil
	}
	if count > n {
		prev.Next = nthFromEnd.Next
		nthFromEnd.Next = nil

	}

	return head
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func checkSl() {
	var sl []byte
	fmt.Println(len(sl), cap(sl))
}

func main() {
	checkSl()

}
