package main

import "fmt"



// 罗马数字有如下符号：

// 基本字符	I	V	X	L	C	D	M
// 对应阿拉伯数字	1	5	10	50	100	500	1000
// 计数规则：
// 相同的数字连写，所表示的数等于这些数字相加得到的数，例如：III = 3
// 小的数字在大的数字右边，所表示的数等于这些数字相加得到的数，例如：VIII = 8
// 小的数字，限于（I、X和C）在大的数字左边，所表示的数等于大数减去小数所得的数，例如：IV = 4
// 正常使用时，连续的数字重复不得超过三次
// 在一个数的上面画横线，表示这个数扩大1000倍（本题只考虑3999以内的数，所以用不到这条规则）

func romanToInt(s string) int {
    m := map[byte] int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }
    fmt.Println(m)
    n := len(s)
    sum := m[s[n-1]]
    for i:=n-2; i>=0; i-- {
        if m[s[i]] < m[s[i+1]] {
            // 小数在大数左边
            sum -= m[s[i]]
        } else {
            sum += m[s[i]]
        }
    }
    return sum
}

func main() {
   fmt.Println(romanToInt("VIII"))
   
}
