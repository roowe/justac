package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    reader := bufio.NewReader(os.Stdin)        
    var n, a int
    var c25, c50, c100 int
    fmt.Fscan(reader, &n)
    var ok = true
    for i:=0; i<n; i++ {
        fmt.Fscan(reader, &a)
        if !ok {
            continue
        }
        if a == 25 {
            c25 += 1
        } else if a == 50 {
            if c25 > 0 {
                c50 += 1
                c25 -= 1
            } else {
                ok = false
            }
        } else {
            if c50 > 0 && c25 > 0{
                c100 += 1
                c50 -= 1
                c25 -= 1
            } else if c25 > 2{
                c100 += 1
                c25 -= 3
            } else {
                ok = false
            }
        }
    }
    if ! ok {
        fmt.Printf("NO")
    } else {
        fmt.Println("YES")
    }
}
