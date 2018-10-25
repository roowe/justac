package main

import (
    "github.com/davecgh/go-spew/spew"
)

type ListNode struct {
    Val  int
    Next *ListNode
}/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
    m := map[int]bool{}
    for _, v := range G {
        m[v] = true
    }
    c := 0
    for ; head != nil; head = head.Next {
        if m[head.Val] && (head.Next == nil || !m[head.Next.Val]) {
            c++
        }
    }
    return c
}
func main() {
    root := &ListNode{Val:1}
    cur := root
    for i:=2; i < 5; i++ {
        n := &ListNode{Val:i}
        cur.Next = n
        cur = n
    }
    spew.Dump(root)
    rev := numComponents(root, []int{0,3,1,4})
    spew.Dump(rev)
}
