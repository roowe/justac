package main

import "fmt"

func abs(n int) int64{
    if n > 0 {
        return int64(n)
    }
    return int64(-n)
}

func divide(dividend int, divisor int) int {
    pmax := 0x7fffffff
    nmax := -pmax - 1
    if dividend == nmax && divisor == -1{
        return pmax
    }

    a, b := abs(dividend), abs(divisor)
    q := 0
    for a >= b {
        // fmt.Println(a, b)
        var i  uint32 = 0
        for ;b << i <= a; i++ {
        }
        //fmt.Println(a, b, i)
        i--
        a -= b << i
        q += 1<<i
    }
    if (dividend>0) == (divisor>0) {
        return q
    }
    return -q
}

func main() {
    fmt.Println(divide(10, 3))
    fmt.Println(divide(7, -3))
    fmt.Println(divide(1, 1))
    fmt.Println(divide(-2147483648, -1))
}
