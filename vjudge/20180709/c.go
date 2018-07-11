package main

import (
    //"math"
    "fmt"
    //"io/ioutil"
)
type Line struct {
    x int
    y int
    z int
}
func main() {
    var n int
    var x0, y0, x, y int
    m := map[Line]bool{}
    fmt.Scan(&n, &x0, &y0)
    for i:=0; i<n; i++ {
        fmt.Scan(&x, &y)
        xd := x0-x
        yd := y0-y
        v := gcd(xd, yd)
        xd /= v
        yd /= v
        z := xd*y0+yd*x0
        m[Line{xd,yd,z}] = true
    }
    fmt.Print(len(m))
}

func gcd(a int, b int) int {
    for b != 0 {
		a, b = b, a%b
	}
    return a
}
