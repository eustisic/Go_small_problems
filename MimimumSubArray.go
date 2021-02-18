/* given an array of positive integers return the minimal length of continguous sub array for which the sum is greater than or equal to a target

ex
target = 7, nums = [2, 3, 1, 2, 4, 3]
output 2 i.e. [4, 3]

ex 
target = 4, num = [1, 4, 4]
output 1 i.e. [4]

ex
target = 11, nums = [1, 1, 1, 1]
output = 0


Qs empty array?
return 0

mental model brute force

find all permutations

nest for loops for each element iterate through each other element and append a slice to a new collection

filter for permutations that equal or are greter than target

filter each slice for a slice that reduces to target or greater than target
 - need to use a reduce method here

return minimum length || sort by length and return length of pos 0

- map the length of each and return the minimum length
math.Min(Infinity, mapped...)
  or
 sort and return value at pos [0]


mental model 

use two pointer that form a sliding scale 
since the numbers are all positive increasing the width of the window will increase the value and decreasing the width of the window will decrease the value

start with the window equal to 0, 0
start with minLength = len of array

set the sum to 1 and the length to r - a + 1 {sum 1, leng 1, tar 4}
[1, 4, 4]
 a
 r

target is greater than sum so need to increase length
[1, 4, 4] {sum 5, len 2, tar 4}
 a
    r
since sum is greater than target commit length to minimum length

increment the anchor since the value of sum is greater than target increment the anchor and subtract the old value from the sum

[1, 4, 4]
    a
    r

sum is equal to 4

*/
package main

import (
	"fmt"
	"math"
)




func minimumSubArray(target int, arr []int) int {
	anchor, reader := 0, 0
	length := +Inf
	sum := arr[0]

	// loop that ends when anchor reaches the end

	for anchor < len(arr) {
		if reader < len(arr) - 1 && sum < target {
			reader++
			sum += arr[reader]
		}

		if sum >= target {
			// modify length since this is a valid substring
			if length > (reader - anchor + 1) {
				length = reader - anchor + 1
			}

			sum -= arr[anchor]
			anchor++
		}

		if sum < target && reader == len(arr) - 1 {
			break
		}
	}

	if length == +Inf {
		return 0
	}

	return length
}

func main() {
	fmt.Println(minimumSubArray(7, []int{1, 2, 3, 0, 4, 3}))
	fmt.Println(minimumSubArray(4, []int{1, 4, 4}))
	fmt.Println(minimumSubArray(10, []int{1, 4, 4}))
}
