package main

import "fmt"


func checkPossibility(nums []int) bool {
    n := len(nums)
    var modified bool
    for i := 1; i < n; i++ {
        if nums[i-1] > nums[i] {
           if modified {
               return false
           }
           if i - 2 < 0 || nums[i - 2] <= nums[i] {
               // i-2, i-1, i
               nums[i - 1] = nums[i]
           } else {
               nums[i] = nums[i - 1]
           }
            modified = true
        }
    }
    //fmt.Println(nums)
    return true
}

func main() {
    fmt.Println(checkPossibility([]int{4, 2, 1}))
    fmt.Println(checkPossibility([]int{4, 2, 3}))
    fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
    fmt.Println(checkPossibility([]int{1, 4, 2, 1, 13}))
}
