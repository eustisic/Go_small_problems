/*

initialize four variable

t = the node at position m - 1
must traverse through the list until node m - 1 is reached
nt = node at m
p = node at m
c = node at m +1
n - node at m + 2

set p.Next = nil

start a for loop starting at 0 while < n - m increment by 1

set c.Next = p
p = c
c = n
if n != nil { n = n.Next }

after loop

t.Next = p
nt.Next = c
*/

func reverseLinkedList(head *ListNode, m int, n int) *ListNode {
	dummy = &ListNode{}
	dummy.Next = head
	t := dummy
	
	for i := 0; i < m - 1; i++ {
		t = t.Next
	}

	nt := t.Next
	p := nt
	c := p.Next
	n := c.Next
	p.Next = nil // break first element of new subList

	for i := 0; i < n - m; i++ {
		c.Next = p
		p = c
		c = n
		if n != nil {n = n.Next}
	}

	t.Next = p
	nt.Next = c

	return head
}
