package main

import "fmt"

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func findDuplicates(nums []int) []int {
    var ret []int

    for _, v := range nums {
        if nums[abs(v)-1] >= 0 {
            nums[abs(v)-1] = -nums[abs(v)-1]
        } else {
            ret = append(ret, abs(v))
        }
    }
    return ret
}

func main() {
    fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1}))
}
