package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

type List struct {
    root *ListNode
}

func newList(val int) *List {
  // takes a value and returns a new list with that value as the root node
  n := ListNode{Val: val}
  l := List{root: &n}
  return &l
}

func (l List) addToList(val int) {
	n := l.root
	for n.Next != nil {
		n = n.Next
	}

	n.Next = &ListNode{Val: val}
}

func (l List) read() {
  n := l.root

  for n != nil {
    fmt.Printf("%d ", n.Val)
    n = n.Next
  }
}

func (l List) readBack() {
  n := l.root
  str := ""

  for n != nil {
    str = fmt.Sprintf("%d ", n.Val) + str    
    n = n.Next 
  }

  fmt.Println(str)
}

func addTwoNumber(l1 *List, l2 *List) *List {
	remainder := 0

	n1 := l1.root
	n2 := l2.root

	firstVal := n1.Val + n2.Val
	remainder = firstVal / 10

	l3 := newList(firstVal % 10)

	for remainder > 0 || n1.Next != nil || n2.Next != nil {
		nextVal := 0
		nextVal += remainder

		if n1.Next != nil {
			nextVal += n1.Next.Val
			n1 = n1.Next
		}

		if n2.Next != nil {
			nextVal += n2.Next.Val
			n2 = n2.Next
		}

		l3.addToList(nextVal % 10)
		remainder = nextVal / 10
	}

	return l3
}


func main() {
	l1 := newList(2)
	l1.addToList(4)
	l1.addToList(3)
	l1.addToList(9)
	l1.addToList(9)

	l2 := newList(5)
	l2.addToList(6)
	l2.addToList(4)
	l2.addToList(1)

  fmt.Println(l1.root)
  fmt.Println(l2.root)

	l3 := addTwoNumber(l1, l2)
	l3.read()
	fmt.Println()
	l3.readBack()
}