/*
PROBLEM
--------

A conveyor belt has packages that must be shipped from one port to another within D days.

The ith package on the conveyor belt has a weight of weights[i]. Each day, we load the ship with packages on the conveyor belt (in the order given by weights). We may not load more weight than the maximum weight capacity of the ship.

Return the least weight capacity of the ship that will result in all the packages on the conveyor belt being shipped within D days.


MENTAL MODELS
------------

Divide and conquer?

[1, 2, 3, 4, 5, 6, 7, 8, 9, 10] 

Sum all the packages and find a mid point based on weight - sum = 55, mid = 27
Divide the packages at the sum mid point

[1, 2, 3, 4, 5, 6, 7] [8, 9, 10]
        28               27

repeat by dividing the largest weight sub-set of packages until you have D (days) separate
sub arrays

[1, 2, 3, 4, 5] [6, 7] [8, 9, 10]
     15           13       27
                 |
                 V

[1, 2, 3, 4, 5] [6, 7] [8, 9] [10]
     15           13      17   10
                 |
                 V
[1, 2, 3, 4, 5] [6, 7] [8] [9] [10]  
     15           13     8  9   10        


maxCapacity = 15, Correct answer, yayyyy!
but... hard to code, time and memory intensive, and not really a binary search 


Binary Search

If we shipped it all in one day would need a capacity equal to the sum of weights

Max amount of days would be a capacity equal to the greatest weight package

[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

minCapacity - 10
maxCapacity - 55

lo = 10
hi = 55
mid = 32

perform a binary search on the interval 10 - 55 (mid point starts at 32)

how many days will it take to ship packages if the sum of shipped packages on a given day can't be greter than our capcity?

iterate through packages adding a day everytime the sum of packages is greater than capacity

(1, 2, 3, 4, 5, 6, 7) (8, 9, 10)
    day 1                day 2

compare the days at searched capacity to required days:

2 days < 5 days so our capacity is too great and we need to search a small capacity

lo = 10, 
hi = 32, 
mid = 21

keep searhing until lo meets hi



*/

package main

import "fmt"

func sumAndMax(arr []int) (int, int) {
    max, sum := 0, 0
    
    for _, n := range arr {
        if max == 0 || n > max {
            max = n
        }
        sum += n
    }
    return max, sum
}

func shipWithinDays(weights []int, D int) int {
    lo, hi := sumAndMax(weights)
    
    for lo < hi {
        days := 1
        sum := 0
        mid := hi + ((lo - hi) >> 1) // capacity to test
        
        // extract into separate function 
        for _, w := range weights {
            
            if sum + w > mid {
                days += 1
                sum = w
            } else {
                sum += w
            }
        }
        
        
        if days <= D {
            hi = mid
        } else {
            lo = mid + 1
        }  
    }
    return lo
}

func main() {
    fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}