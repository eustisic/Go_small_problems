 package main

 import "fmt"

 type ListNode struct {
	Val int
	Next *ListNode
}

 type Queue struct {
	Head *ListNode
	Tail *ListNode
	Length int
}

func (q *Queue) enqueue(l *ListNode) bool {
	if q.Length == 0 || q.Tail == nil {
		q.Head = l
		q.Tail = l
	} else {
		q.Tail.Next = l
		q.Tail = l
	}

	q.Length++
	return true
}

func (q *Queue) dequeue() (dQd int, ok bool) {

	if q.Length == 0 || q.Head == nil {
		ok = false

	} else {
		dQd = q.Head.Val
		ok = true
		q.Head = q.Head.Next
		q.Length--
	}

	if !ok {
		panic("Don't throw away your shot!")
	}

	return dQd, ok
}

func main() {
	q := Queue{}
	n := ListNode{Val: 1}
	fmt.Println(q.enqueue(&n))
	fmt.Println(q.dequeue())
}