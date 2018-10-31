package main

import (
    "fmt"
    "sort"
)

func threeSum(nums []int) [][]int {
    n := len(nums)

    sort.Ints(nums)

    //fmt.Println(nums)
    var ret [][]int
    for i:=0; i<n; i++ {
        // 跳过处理的
        if i != 0 && nums[i] == nums[i - 1]{
            continue
        }
        j := i+1
        k := n-1
        for j < k{
            //fmt.Println(nums[i], nums[j], nums[k])
            s := nums[i] + nums[j] + nums[k]
            if s == 0 {
                ret = append(ret, []int{nums[i], nums[j], nums[k]})
                j++
                for j < k && nums[j] == nums[j-1]{
                    j++
                }
            } else if s < 0 {
                j++
            } else {
                k--
            }
        }
    }
    //fmt.Println(m)
    return ret
}

func main() {
    fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
    //fmt.Println(threeSum(data1()))
}

func data1() []int {
    return []int{9,-9,4,12,12,0,-14,-7,10,-1,11,-10,-3,2,-9,0,8,-9,-5,-1,10,5,13,-5,-9,-12,9,-3,10,10,-10,4,8,1,-7,-2,-14,-6,6,11,8,-6,9,13,11,7,-10,-4,14,0,3,1,-10,-7,3,-12,-3,-11,0,-8,-15,5,3,8,13,11,13,-15,-9,4,3,6,5,-11,-4,-6,4,1,5,-5,-13,-7,11,-8,2,-1,-12,14,3,3,13,-5,-14,-7,11,14,-11,9,6,-13,-9,-13,1,11,-9,12,-10,2,-1,3,12,-7,3,0,0,12,6,3,3,-13,14,1,-3}
}