package main

import "fmt"

type Sortable struct {
  arr []int
}

func newSortable(arr []int) *Sortable {
  s := Sortable{arr: arr}
  return &s
}

func (s *Sortable) partition(p_left, p_right int) int {
	pivot_i := p_right
	pivot := s.arr[pivot_i]
	p_right--
	
	for {
	
	  for s.arr[p_left] < pivot {
	    p_left++
	  }
	
	  for s.arr[p_right] > pivot {
	    if p_right == 0 {
	      break
	    }
	    p_right--
	  }
	
	  if p_left >= p_right {
	    break
	  } else {
	    s.arr[p_left], s.arr[p_right] = s.arr[p_right], s.arr[p_left]
	    p_left++
	  }
	}
	
	s.arr[p_left], s.arr[pivot_i] = s.arr[pivot_i], s.arr[p_left]
	return p_left
}

func (s *Sortable) quickSort(left, right int) {
  if right - left <= 0 {
    return
  }

  pivot_i := s.partition(left, right)

  s.quickSort(left, pivot_i - 1)
  s.quickSort(pivot_i + 1, right)
}

func main() {
  arrSort := newSortable([]int{0, 5, 2, 1, 6, 3, 8})
  arrSort.quickSort(0, len(arrSort.arr) - 1)
  fmt.Println(arrSort.arr)
}
