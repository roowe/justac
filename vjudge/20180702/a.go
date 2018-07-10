package main

import (
    //"math"
    "fmt"
)

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    if a == 0 && b == 0 {
        fmt.Printf("NO")
    } else if ( a == b || a-b == 1 || a - b == -1) {
        fmt.Printf("YES")
    } else {
        fmt.Printf("NO")
    }
}
