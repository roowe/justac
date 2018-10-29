package main

import "fmt"

func transpose(A [][]int) [][]int {
    // ^A0 <---x--->
    // |
    // |
    // y
    // |
    // |
    // V

    x := len(A[0])
    y := len(A)
    B := make([][]int, x)
    for i:=0; i<x; i++ {
        B[i] = make([]int, y)
    }
    for i:=0; i<y; i++ {
        for j:=0; j<x; j++ {
            B[j][i] = A[i][j]
        }
    }
    return B
}

func main() {
    A := make([][]int, 3)
    A[0] = []int{1,1,0}
    A[1] = []int{1,0,1}
    A[2] = []int{0,0,0}
    fmt.Println(transpose(A))
}