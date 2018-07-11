package main

import (
    //"math"
    "fmt"
    //"io/ioutil"
)

func main() {
    // for i:=0; i<=10; i++ {
    //     calc(i)
    // }
    var s string
    fmt.Scan(&s)
    n := len(s)
    var sum int
    if n == 1 {
        sum = int(s[0]-'0')
    } else {
        sum = 2*int(s[len(s)-2]-'0') + int(s[len(s)-1] - '0')
    }
    sum = sum%4
    if sum == 0 {
        fmt.Print(4)
    } else {
        fmt.Print(0)
    }
}
// 0 = 4
// 1 = 0
// 2 = 0
// 3 = 0
// 4 = 4
// 5 = 0
// 6 = 0
// 7 = 0
// 8 = 4
// 9 = 0
// 10 = 0
func calc(n int) {
    var sum int
    for i:=1; i<=4; i++ {
        var p = 1
        for j:=1; j<=n; j++ {
            //fmt.Println(j, i, sum, p)
            p = (p*i) % 5
        }
        sum += p%5
    }
    fmt.Printf("%d = %d\n", n, sum%5)
}
