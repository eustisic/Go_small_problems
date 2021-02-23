Remove Dups from Sorted List II

Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

Starting at the head of the list initialize a slow pointer and a fast pointer

1->2->3->3->4->4->5
s
   f

1. if the slow pointer value is not equal to the fast pointer value increment both s = s.Next f = f.Next

1->2->3->3->4->4->5
      s
         f

2. when the slow pointer and the fast pointer share the same value move the fast pointer until the slow pointer value does not equal the fast pointer value

1->2->3->3->4->4->5
      s
            f

3. set the slow pointer value to the fast pointer value and the slow pointer s.Next to f.Next and f to s.Next

1->2->4->4->5
      s  
         f

2. when the slow pointer and the fast pointer share the same value move the fast pointer until the slow pointer value does not equal the fast pointer value

1->2->4->4->5
      s  
            f

3. set the slow pointer value to the fast pointer value and the slow pointer s.Next to f.Next and f to s.Next

1->2->5->nil
      s  
          f

func spliceList(sPt, fPt *ListNode) {
	sPt.Val = fPt.Val
	sPt.Next = fPt.Next
	fPt = sPt.Next
}         

func removeDups(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	sPt := head
	fPt := head.Next

	for fPt != nil {
		if sPt.Val != fPt.Val {
			sPt = sPt.Next
			fPt = fPt.Next
		} else {

			for sPt.Val == fPt.Val {
				fPt = fPt.Next
			}

			spliceList(sPt, fPt)
		}
	}

	return head
}
