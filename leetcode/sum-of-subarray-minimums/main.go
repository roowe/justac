package main

import "fmt"

type stack []int

func (s stack) Empty() bool { return len(s) == 0 }
func (s stack) Peek() int   { return s[len(s)-1] }
func (s *stack) Put(i int)  { *s = append(*s, i) }
func (s *stack) Pop() int {
    d := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return d
}

func sumSubarrayMins(A []int) int {
    s := make(stack, 0, len(A))
    var mod int64 = 1e9 + 7
    var res, value int64

    for i, v := range A {
        if s.Empty() || v >= A[s.Peek()] {
            s.Put(i)
            value = (value + int64(v)) % mod
        } else {
            for !s.Empty() && v < A[s.Peek()] {
                idx := s.Pop()
                if s.Empty() {
                    value = (value - (int64(idx)+1)*int64(A[idx])) % mod
                } else {
                    value = (value - (int64(idx-s.Peek()))*int64(A[idx])) % mod
                }
            }
            if s.Empty() {
                value = (value + (int64(i)+1)*int64(v)) % mod
            } else {
                value = (value + (int64(i-s.Peek()))*int64(v)) % mod
            }
            s.Put(i)
        }
        res = (value + res) % mod
    }
    return int(res)
}

func main() {
    fmt.Println(sumSubarrayMins([]int{3,1,2,4}))
    fmt.Println(sumSubarrayMins([]int{48,87,27}))
}