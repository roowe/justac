package main

import "fmt"
// 用two point，头尾指针也可以
func twoSum(nums []int, target int) []int {
    for i, v := range nums {
        if ii := searchLast(nums, target-v); ii != -1 {
            return []int{1+i, 1+ii}
        }
    }
    return []int{-1,-1}
}

func searchLast(nums []int, target int) int {
    low, high := 0, len(nums)-1
    for low <= high {
        m := (low+high)/2
        if target > nums[m] {
            // 右边
            low = m+1
        } else if target < nums[m] {
            // 左边
            high = m-1
        } else {
            if m == len(nums)-1 || nums[m+1] != target {
                return m
            }
            low = m+1
        }
    }
    return -1
}

func main() {
    fmt.Println(twoSum([]int{2,7,11,15}, 9))
}
