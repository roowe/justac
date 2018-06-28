package main

import (
    //"math"
    "fmt"
)

func main() {
    var t, s, x int
    fmt.Scan(&t, &s, &x)
    if x < t || t+1 == x{
        no()
    } else if x == t {
        yes()
    } else {        
        r := (x-t)%s
        if r == 1 || r == 0 {
            yes()
        } else {
            no()
        }
    }       
}

func yes() {
    fmt.Printf("YES")
}
func no() {
    fmt.Printf("NO")
}
