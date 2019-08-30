package main

import (
    "fmt"
    "math"

    "github.com/kavehmz/prime"
)

func main() {
    // ans: 1739023853137
    var limit uint64 = 100000000
    ps := prime.Primes(limit)
    m := make(map[uint64]struct{}, len(ps))
    for _, p := range ps {
        m[p] = struct{}{}
    }
    var sum uint64
    var candidate, j uint64
    for _, p := range ps {
        candidate = p - 1
        ok := true
        sq := uint64(math.Sqrt(float64(candidate)))
        for j = 2; j <= sq; j++ {
            if candidate%j == 0 {
                if _, found := m[candidate/j+j]; !found {
                    ok = false
                    break
                }
            }
        }
        if ok {
            //fmt.Println(candidate)
            sum += candidate
        }
    }
    fmt.Println("ans:", sum)
}
