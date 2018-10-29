package main

import "fmt"

func flipAndInvertImage2(A [][]int) [][]int {
    for i:=0; i<len(A); i++ {
        size := len(A[i])
        tmp := make([]int, size)
        copy(tmp, A[i])
        for j:=0; j<size; j++ {
            A[i][j] = 1 - tmp[size-1-j]
        }
    }
    return A
}

func flipAndInvertImage(A [][]int) [][]int {
    for i:=0; i<len(A); i++ {
        size := len(A[i])
        for j:=0; j<(size+1)/2; j++ {
            tmp := A[i][j] ^ 1
            A[i][j] = A[i][size-1-j] ^ 1
            A[i][size-1-j] = tmp
        }
    }
    return A
}
func main() {
    A := make([][]int, 3)
    A[0] = []int{1,1,0}
    A[1] = []int{1,0,1}
    A[2] = []int{0,0,0}
    fmt.Println(flipAndInvertImage(A))
}

