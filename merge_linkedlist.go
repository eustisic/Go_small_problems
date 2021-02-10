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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var root *ListNode
	var l3 *ListNode
	var lowVal int

	for l1 != nil || l2 != nil {
		fmt.Println(l1, l2)
		
		if l1 != nil && l2 != nil {
		  if l1.Val <= l2.Val {
		    lowVal = l1.Val
		    l1 = l1.Next
		  } else {
		    lowVal = l2.Val
		    l2 = l2.Next
		  }
		}
		
		if l1 == nil {
		  lowVal = l2.Val
		  l2 = l2.Next
		} else if l2 == nil {
		  lowVal = l1.Val
		  l1 = l1.Next
		}
			

		if root == nil {
			l3 = &ListNode{Val: lowVal}
			root = l3
			fmt.Println(root)
		} else {
			l3.Next = &ListNode{Val: lowVal}
			l3 = l3.Next
		}
	}

	return root
}

func main() {
	l1 := createList([]int{1,2,4})
	l2 := createList([]int{1,3,4})

	readList(l1)
	fmt.Println()
	readList(l2)
	
	fmt.Println()
	lr := mergeTwoLists(l1, l2)
	readList(lr)
}