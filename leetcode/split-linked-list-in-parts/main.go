package main

import (
    "github.com/davecgh/go-spew/spew"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

func lenList(l *ListNode) int {
    var n int
    for i := l; i != nil; i = i.Next {
        n++
    }
    return n
}

func splitListToParts(root *ListNode, k int) []*ListNode {
    res := make([]*ListNode, k)
    n := lenList(root)
    c := n / k
    r := n % k
    cur := root
    //fmt.Println(n, c, r, k)
    for i:=0; i<k; i++ {
        start := cur
        var j = 1
        if i < r {
            j = 0
        }
        for ; j<c; j++ {
            cur = cur.Next
        }
        //fmt.Println(i, r)

        if cur != nil {
            next := cur.Next
            cur.Next = nil
            cur = next
        }
        res[i] = start
    }
    return res
}

func main() {
    root := &ListNode{Val:1}
    cur := root
    for i:=2; i < 11; i++ {
        n := &ListNode{Val:i}
        cur.Next = n
        cur = n
    }
    spew.Dump(root)
    rev := splitListToParts(root, 3)
    spew.Dump(rev)
}
