// Merge two sorted linked lists

// Algorithm:

// Create a dummy head, an empty node that references a new linked lists &LinkedList

// iterate through the two linked lists keeping track of the dummy head, the current list node and the current node of each linked list

// if the value at l1 is greater than the value at l2 the
// next node in the merged linked list will be l2 and the current node in the linked list is now l2

// - p := lesser of two nodes


// else do the opposite

// if either l1 or l2 is null exit the loop


// if l1 is not null attach the rest of l1 to the end of the merged list
// if l2 is not null attach the rest of l2 to the end fo the list

// return the dummyHead.next

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    p := dummy

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			p.Next = l2
			l2 = l2.Next
		} else {
			p.Next = l1
			l1 = l1.Next
		}
		p = p.Next
	}

	if l1 != nil { p.Next = l1 }
	if l2 != nil { p.Next = l2 }

	return dummy.Next
}