package main

import "fmt"

func numJewelsInStones(J string, S string) int {
    c := make(map[rune]int)
    for _, char := range(S) {
        c[char] += 1
    }
    var sum int = 0
    for _, char := range(J) {
        sum += c[char]
    }
    return sum
}

func main() {
    J := "aA"
    S := "aAAbbbb"
    ret := numJewelsInStones(J, S)
    fmt.Printf("%v\n", ret)
}
