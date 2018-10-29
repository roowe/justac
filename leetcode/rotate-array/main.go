package main

import (
    "fmt"
)

func rotate2(nums []int, k int) {
    k = k % len(nums)
    copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}

func rotate(nums []int, k int) {
    k = k % len(nums)
    reverse(nums[:len(nums)-k])
    reverse(nums[len(nums)-k:])
    reverse(nums)
}
func reverse(a []int)  {
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7}
    rotate(nums, 3)
    fmt.Println(nums)
}
