func search(array []int, target int) int {
    left, right := 0, len(array)-1
    for left+1 < right {
      mid := left + ((right - left) / 2)
      if array[mid] == target {
        return mid
      } else if array[mid] > array[left] {
          if array[left] < target {
            right = mid
          } else {
            left = mid
          }
      } else if array[mid] < array[left] {
          if array[mid] > target {
              right = mid
          } else {
              left = mid
          }
      }
    }

    // loop ends when left and right pointers are next to each other; either could be the result

    if array[left] == target {
      return left
    }
    if array[right] == target {
      return right
    }
    return -1
}
/*
[4,5,6,7,0,1,2]
         se
 
 m is greater than s?
 s is less than target? - no s= m+1
 is m greater than s - yes
 if s is less than or equal to target? - e m - 1


if the target is less than the mid point  but greater than the start
then the value could be there and we perform a binary search as normal

4 cases
target is greater than mid point less than end - happy search as normal
target is less than mid point greater than end - happy search as normal
target is greater than m greater but less e - sad search oposite
target is less than s and less than m - sad search oposite

*/
