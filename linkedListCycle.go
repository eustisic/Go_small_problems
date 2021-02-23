

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 /*
create slow and fast pointers the slow pointer will move one node at a time the fast pointer will move two at a time
if the fast is ever less than the slow then it is a circular linked list.  if the fast is ever nil it is not circular

how to determine if the fast node has passed the slow node? 
can wait until they equal each other (i.e. the address of slow == address of fast) could take a long time if the list is long
can have slow check around it for fast (i.e) slow -1 and slow + 1
*/

func hasCycle(head *ListNode) bool {
		if head == nil || head.Next == nil {
			return false
		}

		slowPt := head.Next
    fastPt := head.Next.Next

		for fastPt != nil && fastPt.Next != nil {
			if fastPt == slowPt {
				return true
			}
			fastPt = fastPt.Next.Next
			slowPt = slowPt.Next
		}

		return false
}