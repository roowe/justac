package main

import "fmt"

func binSearch(nums []int, target, low, high int) int {
    m := (low + high) / 2
    if low > high {
        return -1
    }
    if nums[low] > nums[high] {
        // 反序
        if i := binSearch(nums, target, low, m); i != -1 {
            return i
        }
        if i := binSearch(nums, target, m+1, high); i != -1 {
            return i
        }
    } else {
        for low <= high {
           m := (low + high) / 2
           if target == nums[m] {
               return m
           } else if target > nums[m] {
               // 右边
               low = m + 1
           } else {
               // 左边
               high = m - 1
           }
        }
        return -1
        // if target == nums[m] {
        //    return m
        //} else if target > nums[m] {
        //    // 右边
        //    low = m + 1
        //} else {
        //    // 左边
        //    high = m - 1
        //}
        //return binSearch(nums, target, low, high)
    }
    return -1
}
func search(nums []int, target int) int {
    n := len(nums)
    low, high := 0, n-1
    return binSearch(nums, target, low, high)
}

func main() {
    fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
    fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
    fmt.Println(search([]int{3, 5, 1}, 3))
    fmt.Println(search([]int{3, 1}, 3))

}
