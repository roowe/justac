package main

import "fmt"
// 10^k % 9 = 1 
// a*10^k % 9 = a % 9
// 假如有一个数字 x = 23456
// x = 2* 10000 + 3 * 1000 + 4 * 100 + 5 * 10 + 6
// 2 * 10000 % 9 = 2 % 9
// 3 * 1000 % 9 = 3 % 9
// 4 * 100 % 9 = 4 % 9
// 5 * 10 % 9 = 5 % 9
// 那么x % 9 = ( 2+ 3 + 4 + 5 + 6) % 9, 注意 x = 2* 10000 + 3 * 1000 + 4 * 100 + 5 * 10 + 6
// 所以有： 23456 % 9 = (2 + 3 + 4 + 5 + 6) % 9
func addDigits(num int) int {
    return (num-1)%9+1
}
func main() {
    // fmt.Printf("%v\n", parse(342))
    // fmt.Printf("%v\n", parse(465))
    fmt.Printf("%v\n", addDigits(0))
    fmt.Printf("%v\n", addDigits(38))
}
