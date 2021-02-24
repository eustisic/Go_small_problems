[2, 3, 1, 1, 4]
    ^

are we at the end - yes the end is true

cache{4: true}

from idx 3 with a maximum range of 1 can we reach a true value in the cache
cache[3 + 1]? it is true therefore: cache{4: true, 3: true}

from idx 2 with a max range of 1 can we reach a true value in the cache?
cache{4: true, 3: true, 2: true}

from idx 1 with max range of 3 can we reach a true value in the cache?

 cache{4: true, 3: true, 2: true, 1: true}


return cache[0]






// func canJump(nums []int) bool {
//     cache := map[int]bool{}
//     return helper(nums, 0, cache)
// }

// func helper(nums []int, idx int, cache map[int]bool) bool {
//     if idx + 1 == len(nums) {
//         return true
//     }
    
//     _, ok := cache[idx]
    
//     if !ok {
//         for i := nums[idx]; i > 0; i-- {
//             if i + idx < len(nums) && !cache[idx] {
//                 cache[idx] = helper(nums, idx + i, cache)
//             } 
//         }
//     }

//     return cache[idx]
// }

