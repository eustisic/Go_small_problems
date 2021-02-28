func permute(nums []int) [][]int {
  var results [][]int
  var temp []int
  backtrack(nums, temp, &results)
  return results
}

// temp     -> current path of exploration
// results  -> successful finds

func backtrack(list []int, temp []int, results *[][]int) {
  if len(temp) > len(list) {
    return
  }

  if len(temp) == len(list) {
    copyTemp := make([]int, len(temp))
    copy(copyTemp, temp)
    *results = append(*results, copyTemp)
    return
  }

  for i := 0; i < len(list); i++ {
    if includes(temp, list[i]) { 
      continue
    }

    temp = append(temp, list[i])     // take
    backtrack(list, temp, results) // explore
    temp = temp[:len(temp) - 1]     // clean up
  }
}

func includes(arr []int, val int) bool {
    for _, v := range arr {
        if v == val {
            return true
        }
    }
    return false
}