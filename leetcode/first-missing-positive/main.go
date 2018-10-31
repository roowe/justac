package main

import "fmt"

func firstMissingPositive(nums []int) int {
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] > 0 && nums[i] <= n && nums[i] != i+1 {
            if nums[nums[i]-1] != nums[i] {
                tmp := nums[nums[i]-1]
                nums[nums[i]-1] = nums[i]
                nums[i] = tmp
                i--
            }
        }
    }
    for i:=0; i<n; i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }
    return n+1
}

func main() {
    fmt.Println(firstMissingPositive([]int{1,2,0}))
    fmt.Println(firstMissingPositive([]int{3,4,-1,1}))
    fmt.Println(firstMissingPositive([]int{7,8,9,11,12}))
}
