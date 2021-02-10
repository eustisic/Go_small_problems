package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func createList(arr []int) *ListNode {
	var firstNode *ListNode
	var prevNode *ListNode

	for i, v := range arr {
		if (i == 0) {
			firstNode = &ListNode{Val: v}
			prevNode = firstNode
		} else {
			prevNode.Next = &ListNode{Val: v}
			prevNode = prevNode.Next
		}
	}

	return firstNode
}

func readList(list *ListNode) {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	Last := head
	var nextLast *ListNode

	for j := k; j > 0; j-- {
		for Last.Next != nil {
			nextLast = Last
			Last = Last.Next
		}

		Last.Next = head
		head = Last
		nextLast.Next = nil
	}

	return head
}

func main() {
	l1 := createList([]int{1,2,3,4,5})
	readList(l1)
	l1 = rotateRight(l1, 2)
	fmt.Println()
	readList(l1)
}