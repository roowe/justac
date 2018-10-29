package main

import (
    "fmt"
    "sort"
)

func arrayPairSum(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    sum := 0
    for i:=0; i<len(nums)/2; i++ {
        sum += nums[2*i]
    }
    return sum
}

func main() {
    fmt.Println(arrayPairSum([]int{1,4,3,2}))
}
