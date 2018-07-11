package main

import (
    //"math"
    "fmt"
    //"io/ioutil"
)

func main() {
    var n int
    var s string
    var d, sum int
    var f bool
    fmt.Scan(&n)
    for i:=0; i<n; i++ {
        fmt.Scan(&d, &s)
        if s[0] != 'S' && sum == 0 {
            f = true
            break
        }
        if s[0] != 'N' && sum == 20000 {
            f = true
            break
        }
        if s[0] == 'S' {
            sum += d
        }
        if s[0] == 'N' {
            sum -= d
        }
        if sum<0 || sum>20000 {
            f = true
            break
        }
    }
    if !f && sum == 0 {
        fmt.Print("YES")
    } else {
        fmt.Print("NO")
    }
}
