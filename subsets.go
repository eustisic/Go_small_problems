func subsets(nums []int) [][]int {
  var results [][]int
  var temp []int
    sort.Ints(nums)
  backtrack(0, nums, temp, &results)
  return results
}



func backtrack(pos int, list []int, temp []int, results *[][]int) {
  if pos > len(list) {
    return
  }

    
    copyTemp := make([]int, len(temp))
    copy(copyTemp, temp)
    *results = append(*results, copyTemp)

  for i := pos; i < len(list); i++ {
      if includesOrLess(temp, list[i]) { 
      continue
    }

    temp = append(temp, list[i])     // take
    backtrack(pos+1, list, temp, results) // explore
    temp = temp[:len(temp) - 1]     // clean up
  }
}

func includesOrLess(temp []int, val int) bool {
    for _, v := range temp {
        if val == v {
            return true
        }
    }
    if len(temp) > 0 {
        return temp[len(temp)-1] > val
    }
    return false
}Í