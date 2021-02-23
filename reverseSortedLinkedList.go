/*
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NIL
Output: 5->4->3->2->1->NIL

use a k sliding scale method
where previos current and next refer to nodes or nil

NIL 1->2->3->4->5->NIL
p   c  n

as long as head is not equal to nil
p := nil
c := head
n := head.Next

Loop


	1. assign current to previous
	c.Next = p

	NIL <- 1  2->3->4->5->NIL
	p      c  n

	2. move p to c, move c to n move n to n.Next
	p = c
	c = n
	if n.Next != nil n = n.Next

	NIL <- 1  2->3->4->5->NIL
	       p  c  n

	if c == nil

	NIL <- 1<-2<-3<-4<-5 NIL 
	                   p  c   n

break out of the loop and return previous



*/

func reverseLinkedList(head *ListNode) *ListNode {
	var p *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	c := head
	n := head.Next

	for c != nil {
		c.Next = p
		p = c
		c = n

		if n != nil { n = n.Next }
	}

	return p
}