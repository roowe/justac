package main

import "github.com/davecgh/go-spew/spew"

type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    fakeHead := &ListNode{Next:head}
    prev := fakeHead
    for head!=nil && head.Next!=nil {
        // prev
        //      head 1 head.Next 2 head.Next.Next 3
        prev.Next = head.Next // 2
        head.Next = head.Next.Next // 3
        prev.Next.Next = head // 1
        // prev
        //      head.Next 2 head 1 head.Next.Next 3
        prev = head
        head = head.Next
    }
    return fakeHead.Next
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
    rev := swapPairs(root)
    spew.Dump(rev)
}
