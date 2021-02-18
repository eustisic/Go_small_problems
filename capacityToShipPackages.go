import "fmt"

func sumAll(weights []int) int {
    sum := 0
    
    for _, n := range weights {
        sum += n
    }
    
    return sum
}

func shipWithinDays(weights []int, D int) int {
    // find the minimum of the wieghts array
    // find the sum of the weights array
    max := sumAll(weights)
    min := 0
    
    for i, n := range weights {
        if i==0 || n > min {
            min = n
        }
    }
    
    for i := min; i < max; i++ {
        res := [][]int{}
        sub := []int{}
        sum := 0

        k := 0

        // break up the array into weights by capacity
        for j := 0; j < len(weights); j++ {
            sum += weights[j]
            sub = append(sub, weights[j])

            if sum > i {
              res = append(res, sub)
              sub = []int{}
            }
        }

        // return the working capacity
        if len(res) == D {
            return i
        }
    }

    return -1
}

