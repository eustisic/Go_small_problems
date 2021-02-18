
package main

import "fmt"


// to mutate the array we can push all elements that are the int to the end of the array using pointers

// have a pointer at each end if the first element
  // move the leading pointer forward when the element is not equal to the val
  // move the trailing pointer forward when the element is equal to the val
 // if the leading pointer is euqal to the val swap with the trailing pointer

// continue until leading == trailing

// O(N) runtime O(1) time complexity

func removeElement(nums []int, val int) int {
    lead := 0
    tail := len(nums) - 1
    
    for tail >= lead {
        for tail >= 0  && nums[tail] == val  {
            tail--
        }
        
        for lead <= tail && nums[lead] != val {
            lead++
        }

        if lead < tail {
            nums[lead], nums[tail] = nums[tail], nums[lead]
        }
        
    }
    
    return lead
}

// alternate solution

// create two pointers starting at 0
// increment one pointer the leading pointer and if is not on the desire value
// swap the desired value with the anchor pointer
// increment the anchor pointer

// O(N) runtime O(1) time complexity



func removeElement(nums []int, val int) int {
    i := 0
    for j := 0; j < len(nums); j++ {
        if nums[j] != val {
            nums[i] = nums[j]
            i++
        }
    }
    return i
}

func main() {
    fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}