package main

import "fmt"



func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    x0 := x
    ret := 0   
    for x != 0{        
        ret = ret*10 + x % 10
        x = x / 10            
    }
    const MIN = -0x80000000
    const MAX = 0x7FFFFFFF
    // fmt.Println(ret)
    return ret == x0    
}

func main() {
   fmt.Println(isPalindrome(123456))
   fmt.Println(isPalindrome(-123456))
   fmt.Println(isPalindrome(-2147447412))
   fmt.Println(isPalindrome(1))
}
