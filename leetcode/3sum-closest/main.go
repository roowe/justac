package main

import (
    "fmt"
    "math"
    "sort"
)

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func threeSumClosest(nums []int, target int) int {
    n := len(nums)
    sort.Ints(nums)
    //fmt.Println(nums)
    var ret int
    var minDiff = math.MaxInt32
    for i := 0; i < n-2; i++ {
        // 跳过处理的
        if i > 0 && nums[i] > 0 && nums[i] > target {
            return ret
        }
        j := i + 1
        k := n - 1
        for j < k {
            //fmt.Println(nums[i], nums[j], nums[k])
            s := nums[i] + nums[j] + nums[k]
            diff := abs(s - target)
            if diff < minDiff {
                ret = s
                minDiff = diff
            }
            if s < target {
                j++
            } else if s >= target {
                k--
            } else {
                return s
            }
        }
    }
    //fmt.Println(m)
    return ret
}

func main() {
    fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
