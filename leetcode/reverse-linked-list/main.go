package main

import (
    "github.com/davecgh/go-spew/spew"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for ; head != nil; {
        //fmt.Printf("head %v\n", head)
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}

func main() {
    root := &ListNode{Val:0}
    cur := root
    for i:=1; i < 5; i++ {
        n := &ListNode{Val:i}
        cur.Next = n
        cur = n
    }
    spew.Dump(root)
    rev := reverseList(root)
    spew.Dump(rev)
}
