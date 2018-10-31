package main

import "fmt"

func removeDuplicates2(nums []int) int {
    n := len(nums)
    var i, j int
    for i, j = 1, 1; i < n; {
        for i<n && nums[i-1] == nums[i] {
            i++
        }
        //fmt.Println("i", i, j)
        if i<n {
            nums[j] = nums[i]
            i, j = i+1, j+1
        }
    }
    //fmt.Println("return i", i, j)
    return j
}

func removeDuplicates(nums []int) int {
    var i int
    if len(nums) > 0 {
        i = 1
    }
    for _, v := range nums {
        if nums[i-1] != v {
            nums[i] = v
            i++
        }
    }
    return i
}
func main() {
    //nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
    nums := []int{1, 2}
    fmt.Println(nums)
    fmt.Println(removeDuplicates(nums))
    fmt.Println(nums)
}
