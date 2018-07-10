package main

import (
    "math"
    "fmt"
)

func main() {
    var r, x, y, x1, y1 int64
    fmt.Scan(&r, &x, &y, &x1, &y1)
    f := math.Sqrt(float64((x-x1)*(x-x1)+(y-y1)*(y-y1)))/float64(r)/2.0
    fmt.Print(math.Ceil(f))
}
