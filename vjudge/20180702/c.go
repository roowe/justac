package main

import (
    "os"
    "fmt"
    "bufio"
    //"sort"
)

func main() {
    reader := bufio.NewReader(os.Stdin)        
    var n int
    fmt.Fscan(reader, &n)
    a := make([]int, n)
    ps := make([]int, n)
    for i:=0; i<n; i++ {
        fmt.Fscan(reader, &a[i])
    }
    for i:=0; i<n; i++ {
        var sum120, bt120 int
        var sum50, bt50 int
        for j:=i-1; j>=0; j-- {
            if a[i] - a[j] >= 1440 {
                break
            }
            if a[i] - a[j] < 1440 {
                sum120 += ps[j]
                bt120 = a[j]
            }
            if a[i] - a[j] < 90 {
                sum50 += ps[j]
                bt50 = a[j]
            }
        }
        // j50 := sort.Search(i-1, fun(j0 int) bool {
        //     return a[j0] > a[i] - 90
        // })
        // j120 := sort.Search(i-1, fun(j0 int) bool {
        //     return a[j0] > a[i] - 1440
        // })
        // if j50 != i {
        //     sum50 +=
        // }
        //fmt.Println(a[i], bt120, sum120, bt50, sum50)
        if (a[i] - bt50 < 90 && sum50 >= 50) ||
            (a[i] - bt120 < 1440 && sum120 >= 120) {
            fmt.Println(0)
        } else {
            v1 := 120 - sum120
            v2 := 50 - sum50
            v3 := 20            
            ps[i] = min(min(v1,v2), v3)
            fmt.Println(ps[i])
        }
    }
}
func min(a int, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}
