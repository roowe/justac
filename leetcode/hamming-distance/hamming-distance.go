package main

import "fmt"

// func hammingDistance(x int, y int) int {
//     v := x^y
//     n := 0
//     for v!=0 {
//         v &=(v-1)
//         n++
//     }
//     return n
// }
func hammingDistance(x int, y int) int {
    v := x^y
    n := 0
    for v!=0 {
        n += v&1        
        v >>=1
    }
    return n
}
func main() {
    J := 1
    S := 4
    ret := hammingDistance(J, S)
    fmt.Printf("%v\n", ret)
}
