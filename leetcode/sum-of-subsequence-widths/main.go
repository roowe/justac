package main

import (
    "fmt"
    "sort"
)

func sumSubseqWidths(A []int) int {
    sort.Ints(A)
    const mod = 1e9 + 7
    var c = 1
    var res = 0
    var n = len(A)
    for i, v := range A {
        res = (res+v*c - A[n-1-i]*c)%mod
        c = (c << 1) % mod
     }
    return res
}

func main() {
    fmt.Println(sumSubseqWidths([]int{2, 1, 3}))
}
