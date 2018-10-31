package main

import (
    "fmt"
    "sort"
)

func fourSum(nums []int, target int) [][]int {
    n := len(nums)

    sort.Ints(nums)

    //fmt.Println(nums)
    var ret [][]int
    for i := 0; i < n; i++ {
        // 跳过处理的
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
        for ii := i + 1; ii < n; ii++ {
            if ii != i+1 && nums[ii] == nums[ii-1] {
                continue
            }
            j := ii + 1
            k := n - 1
            for j < k {
                //fmt.Println(nums[i], nums[ii], nums[j], nums[k])
                s := nums[i] + nums[ii] + nums[j] + nums[k]
                if s == target {
                    ret = append(ret, []int{nums[i], nums[ii], nums[j], nums[k]})
                    j++
                    for j < k && nums[j] == nums[j-1] {
                        j++
                    }
                } else if s < target {
                    j++
                } else {
                    k--
                }
            }
        }
    }
    //fmt.Println(m)
    return ret
}

func main() {
    //fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
    //fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
    fmt.Println(fourSum([]int{-1,0,1,2,-1,-4}, -1))

    //fmt.Println(threeSum(data1()))
}
