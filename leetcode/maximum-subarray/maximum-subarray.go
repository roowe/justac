package main

import "fmt"
func maxSubArray(nums []int) int {
    max := -1 << 31
    sum := 0
    for _, n := range(nums) {
        sum += n
        if sum > max {
            max = sum
        }
        if sum < 0 {
            sum = 0
        }
    }
    return max
}
func main() {
    // fmt.Printf("%v\n", parse(342))
    // fmt.Printf("%v\n", parse(465))
    fmt.Printf("%v\n", maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}
