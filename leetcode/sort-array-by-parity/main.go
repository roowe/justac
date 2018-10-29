package main

import "fmt"

func sortArrayByParity(A []int) []int {
    even := make([]int, 0, len(A))
    odd := make([]int, 0, len(A))
    for _, v := range A {
        if v%2 == 0 {
            even = append(even, v)
        } else {
            odd = append(odd, v)
        }
    }
    return append(even, odd...)
}

func main() {
    fmt.Println(sortArrayByParity([]int{1,2,3,4,5,6,7,8,9,10,5,4,3,2,1}))
}

