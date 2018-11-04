package main

import "fmt"

func findComplement(num int) int {
    s := 1

    for i := num; i != 0; i /= 2 {
        s *= 2
    }
    //fmt.Println("s", s)
    return num ^ (s - 1)
}

func main() {
    fmt.Println(findComplement(5))
    fmt.Println(findComplement(1))
}
