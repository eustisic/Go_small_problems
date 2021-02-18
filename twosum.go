/*
Given an array of integers, find two numbers such that they add up to a specific target number.
The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution.

problem: search an array of numbers for two numbers that add to a third given number
- assumptions, are the arrays sorted or un sorted
- will there be negative numbers
- will there be floats

example: see below

data: data in is an array of integers and an iteger, data out is two different integers
representing the idx + 1 of two separate values in the array
datastructure: hash table or loop?


algorithm 1, unsorted arrays:
1. create a hash table of integer - integer pairs
2. start a loop and for each el of the array add the difference between the search value as the key and the index as the value
3. if the hash key has the key value pair already return the current index and the value from the hash table
4. loop completes without find a value return -1, -1


algorithm 2 sorted array:
1. create a single loop that keep track of values on either end of the array.  
2. If the sum of the two values is too small move left index right, if value is too large move right index left
3. if the values equal search param return two indexes + 1
4 if left index equals right index return -1, -1
*/



package main

import "fmt"

func twoSumUnsorted(arr []int, search int) (idx1, idx2 int) {
	table := make(map[int]int)

	for i, v := range arr {
		_, prs := table[v]

		if prs {
			return table[v] + 1, i + 1
		} else {
			table[search - v] = i
		}
	}
	return -1, -1
}

func twoSumSorted(arr []int, search int) (int, int) {
	i := 0
	j := len(arr) - 1

	for i < j {
		if (arr[i] + arr[j]) < search {
			i++
		} else if (arr[i] + arr[j]) > search {
			j--
		} else {
			return i + 1, j + 1
		}
	}
	return -1, -1
}

 /*
Design and implement a TwoSum class. It should support the following operations: add and find.
add(input) – Add the number input to an internal data structure.
find(value) – Find if there exists any pair of numbers which sum is equal to the value.
For example,
add(1); add(3); add(5); find(4)true; find(7)false
*/

type TwoSum struct {
	data map[int]int
}

func (t *TwoSum) add(val int) {
	v, p := t.data[val]

	if (p) {
		t.data[val] = v + 1
	} else {
		t.data[val] = 1
	}
}

func (t *TwoSum) find(search int) bool {
	ref := map[int]bool{}

	for key, val := range t.data {
		_, prs := ref[key]

		if (key * val == search) {
			return true
		}

		if prs {
			return true
		} else {
			ref[search - key] = true
		}
	}
	return false
}

func main() {
  fmt.Println(twoSumUnsorted([]int{2, 3, 4, 6, 7}, 10)) // 3 and 4
  fmt.Println(twoSumUnsorted([]int{2, 3, 4, 6, 7}, 12)) // -1, -1

  fmt.Println(twoSumSorted([]int{2, 3, 4, 6, 7}, 10)) // 2 and 5
  fmt.Println(twoSumSorted([]int{2, 3, 4, 6, 7}, 12)) // - 1, -1

  list := TwoSum{data: map[int]int{}}
  list.add(5)
  list.add(1)
  list.add(4)
  list.add(4)
  list.add(3)
  fmt.Println(list.data)

  fmt.Println(list.find(4))
  fmt.Println(list.find(10))
  fmt.Println(list.find(8))
}