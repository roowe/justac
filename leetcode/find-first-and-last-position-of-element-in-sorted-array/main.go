package main

import "fmt"

func searchRange(nums []int, target int) []int {
    n := len(nums)
    l, r := 0, n-1
    for l <= r {
        m := (l+r)/2
        if target > nums[m] {
            // 右边
            l = m+1
        } else if target < nums[m] {
            // 左边
            r = m-1
        } else {
            ll, rr := searchFirst(nums, target, l, r), searchLast(nums, target, l, r)
            return []int{ll, rr}
        }
    }
    return []int{-1, -1}
}
func searchLast(nums []int, target int, low int, high int) int {
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
func searchFirst(nums []int, target int, low int, high int) int {
    for low <= high {
        m := (low+high)/2
        if target > nums[m] {
            // 右边
            low = m+1
        } else if target < nums[m] {
            // 左边
            high = m-1
        } else {
            if m == 0 || nums[m-1] != target {
                return m
            }
            high = m-1
        }
    }
    return -1
}

func main()  {
    fmt.Println(searchRange([]int{1}, 1))

    fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
    fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
}
