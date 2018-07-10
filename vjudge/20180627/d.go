package main

import (
    //"strings"
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    //var sb strings.Builder
    var d, k int
    for i:=0; i<n; i++ {
        if 1+i*2 <= n {
            d = 1+i*2
        } else {
            d = n-(1+i*2%n)
        }
        k = (n-d)/2
        for j:=0; j<k; j++ {
            //sb.WriteString("*")
            fmt.Printf("*")
        }
        for j:=0; j<d; j++ {
            //sb.WriteString("D")
            fmt.Printf("D")
        }
        for j:=0; j<k; j++ {
            fmt.Printf("*")
            //sb.WriteString("*")
        }
        fmt.Println("")
        //sb.WriteString("\n")
    }
    //fmt.Printf(sb.String())
}
