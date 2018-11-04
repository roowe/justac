package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
    m := map[int]int{}

    for _, v := range nums1 {
        m[v] = 1
    }
    var ret []int
    for _, v := range nums2 {
        if c, ok := m[v]; ok && c == 1{
            m[v] = c + 1
            ret = append(ret, v)
        }
    }
    return ret
}

func main ()  {
    fmt.Println(intersection([]int{1,2,2,1}, []int{2,2}))
    fmt.Println(intersection([]int{4,9,5}, []int{9,4,9,8,4  }))
}