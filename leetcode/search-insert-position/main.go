package main

import "fmt"

func searchInsert(nums []int, target int) int {
    low, high := 0, len(nums)-1
    for low <= high {
        m := (low + high) / 2
        if nums[m] < target {
            low = m + 1
        } else if nums[m] == target {
            return m
        } else {
            if m == 0 || target > nums[m-1] {
                // find
                return m
            } else {
                high = m - 1
            }
        }
    }
    return low
}

func main()  {
    data := []int{1,3,5,6}
    fmt.Println(searchInsert(data, 5))
    fmt.Println(searchInsert(data, 2))
    fmt.Println(searchInsert(data, 7))
    fmt.Println(searchInsert(data, 0))
    fmt.Println(searchInsert([]int{1,3,5}, 1))

}
