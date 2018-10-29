package main

import "fmt"

func makesquare(nums []int) bool {
    sums := make([]int, 4)
    sum := 0
    for _, n := range nums {
        sum += n
    }
    if sum % 4 != 0 {
        return false
    }
    sum /= 4
    return dfs(nums, 0, sums, sum)
}

func dfs(nums []int, i int, sums []int, sum int) bool {
    if i == len(nums) {
        //fmt.Println(sums)
        if sums[0] == 0 {
            return false
        }
        for j:=1; j<4; j++ {
            if sums[0] != sums[j] {
                return false
            }
        }
        return true
    }
    for j := 0; j < 4; j++ {
        //fmt.Println(i, j)
        sums[j] += nums[i]
        if sums[j] > sum {
            sums[j] -= nums[i]
            continue
        }
        if dfs(nums, i+1, sums, sum) {
            return true
        }
        sums[j] -= nums[i]
    }
    return false
}
func main() {
    fmt.Println(makesquare([]int{1,2,3,4,5,6,7,8,9,10,5,4,3,2,1}))
    fmt.Println(makesquare([]int{1,1,2,2,2}))
    fmt.Println(makesquare([]int{2,2,2,2,2,2}))
    fmt.Println(makesquare([]int{3,3,3,3,4}))
}