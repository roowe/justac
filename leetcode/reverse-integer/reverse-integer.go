package main

import "fmt"



func reverse(x int) int {
    ret := 0   
    for x != 0{        
        ret = ret*10 + x % 10
        x = x / 10            
    }
    const MIN = -0x80000000
    const MAX = 0x7FFFFFFF
    if ret < MIN || ret > MAX {
        return 0
    } else {
        return ret   
    }  
}

func main() {
   fmt.Println(reverse(123456))
   fmt.Println(reverse(-123456))
   fmt.Println(reverse(127000))
}
