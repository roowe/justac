package main

import (
    "container/list"
    "fmt"
)

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func shortestSubarray(A []int, K int) int {
    sum := make([]int, len(A)+1)
    for i, v := range A {
        sum[i+1] = sum[i] + v
    }
    dq := list.New()
    ans := len(A) + 1
    for i:=0; i<=len(A); i++{
        for dq.Len() != 0 {
            f := dq.Front()
            // fmt.Println("f=>", f.Value, "sum", sum)
            if sum[i] - sum[f.Value.(int)] >= K {
                // fmt.Println("ans", ans)
                ans = min(ans, i - f.Value.(int))
                dq.Remove(f)
            } else {
                break
            }
        }
        for dq.Len() != 0 {
            b := dq.Back()
            if sum[i] <= sum[b.Value.(int)] {
                // 如果j<i，并且sum[j]大于等于sum[i]，j已经被运算过
                dq.Remove(b)
            } else {
                break
            }
        }
        dq.PushBack(i)
    }
    if ans == len(A) + 1 {
        return -1
    }
    return ans
}

func main()  {
    fmt.Println(shortestSubarray([]int{1}, 1))
    fmt.Println(shortestSubarray([]int{1,2}, 4))
    fmt.Println(shortestSubarray([]int{2,-1,2}, 3))
}
