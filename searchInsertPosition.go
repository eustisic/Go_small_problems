func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left+1 < right {
      mid := left + ((right - left) / 2)
      if nums[mid] == target {
        return mid
      } else if nums[mid] < target {
        // always include mid in the next iteration
        left = mid
      } else {
        right = mid
      }
    }

// loop ends when left and right pointers are next to each other; either could be the result

    if nums[left] == target {
      return left
    }
    if nums[right] == target {
      return right
    }
    if nums[right] < target {
        return right + 1
    }
    if nums[left] > target {
        return left
    }
    return right
}

// brute force - linear search through the array
// if the target is found return the index of the target
// if a value that is greater than the target is reached return the index of that value

/*
binary search
perform a binary search of the array if the target is found return the index
of value
target = 5
[1, 3, 5, 6]

look at the midpoint if the value is at the midpoint return the index
if the value is greater than the midpoint find the upper half midpoint etc..
if the value is less than the midpoint search the lower half midpoint etc..

if left index and right index are next to each other and the value is not present
return the right index

*/