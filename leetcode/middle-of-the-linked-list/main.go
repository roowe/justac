package main

import (
    "github.com/davecgh/go-spew/spew"
)

type ListNode struct {
    Val  int
    Next *ListNode
}
func middleNode(head *ListNode) *ListNode {
    fast, slow := head, head
    for ; fast.Next != nil && fast.Next.Next != nil ;{
        slow = slow.Next
        fast = fast.Next.Next
    }
    if fast.Next == nil {
        return slow
    }
    return slow.Next
}

func main() {
    root := &ListNode{Val:1}
    cur := root
    for i:=2; i < 6; i++ {
        n := &ListNode{Val:i}
        cur.Next = n
        cur = n
    }
    spew.Dump(root)
    rev := middleNode(root)
    spew.Dump(rev)

}
